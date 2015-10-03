package storage

import "sync"

type inMemoryDB struct {
	m   map[string][]byte
	lck sync.RWMutex
}

func NewInMemoryDB() DB {
	return &inMemoryDB{m: make(map[string][]byte)}
}

func (d *inMemoryDB) Get(key string) ([]byte, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()
	v, ok := d.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

func (d *inMemoryDB) Set(key string, val []byte) error {
	d.lck.Lock()
	defer d.lck.Unlock()
	d.m[key] = val
	return nil
}
