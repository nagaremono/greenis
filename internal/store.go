package internal

import "sync"

var Store *MemStore = &MemStore{
	table: make(map[string]Resp),
}

type MemStore struct {
	mu    sync.RWMutex
	table map[string]Resp
}

func (s *MemStore) Get(key string) (Resp, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.table[key]
	return val, ok
}

func (s *MemStore) Set(key string, val Resp) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.table[key] = val

	return nil
}

func (s *MemStore) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.table, key)

	return nil
}
