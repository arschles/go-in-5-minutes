package episode0

import "sync"

type inMemoryHashTable struct {
	m   map[string][]byte
	lck sync.RWMutex
}

func NewInMemoryHashTable() HashTable {
	return &inMemoryHashTable{m: make(map[string][]byte)}
}

func (i *inMemoryHashTable) Get(key string) ([]byte, error) {
	i.lck.RLock()
	defer i.lck.RUnlock()
	val, ok := i.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

func (i *inMemoryHashTable) Set(key string, val []byte) error {
	i.lck.Lock()
	defer i.lck.Unlock()
	i.m[key] = val
	return nil
}
