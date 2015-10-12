package storage

import "sync"

type inMemoryDB struct {
	m   map[string][]byte
	lck sync.RWMutex
}

// NewInMemoryDB creates a new DB implementation that stores all data in memory.
// All operations are concurrency safe
func NewInMemoryDB() DB {
	return &inMemoryDB{m: make(map[string][]byte)}
}

// Get is the interface implementation
func (d *inMemoryDB) Get(key string) ([]byte, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()
	v, ok := d.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

// Set is the interface implementation
func (d *inMemoryDB) Set(key string, val []byte) error {
	d.lck.Lock()
	defer d.lck.Unlock()
	d.m[key] = val
	return nil
}
