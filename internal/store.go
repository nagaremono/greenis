package internal

var Store *MemStore = &MemStore{
	table: make(map[string]Resp),
}

type MemStore struct {
	table map[string]Resp
}

func (s *MemStore) Get(key string) (Resp, bool) {
	val, ok := s.table[key]
	return val, ok
}

func (s *MemStore) Set(key string, val Resp) error {
	s.table[key] = val

	return nil
}
