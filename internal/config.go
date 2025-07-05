package internal

import (
	"sync"
)

var cfg *RDBConfig = &RDBConfig{}

type RDBConfig struct {
	Dir      string
	Filename string
	mu       sync.RWMutex
}

func (r *RDBConfig) SetDir(d string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Dir = d
}

func (r *RDBConfig) SetFile(f string) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.Filename = f
}
