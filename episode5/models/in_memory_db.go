package models

import "sync"

type inMemDB struct {
	rwm *sync.RWMutex
	m   map[string][]byte
}

func NewInMemoryDB() DB {
	return &inMemDB{
		rwm: &sync.RWMutex{},
		m:   make(map[string][]byte),
	}
}

func (db *inMemDB) Get(key string) ([]byte, error) {
	db.rwm.RLock()
	defer db.rwm.RUnlock()
	val, ok := db.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

func (db *inMemDB) Set(key string, val []byte) error {
	db.rwm.Lock()
	defer db.rwm.Unlock()
	_, ok := db.m[key]
	if !ok {
		return ErrNotFound
	}
	db.m[key] = val
	return nil
}

func (db *inMemDB) Upsert(key string, val []byte) (bool, error) {
	db.rwm.Lock()
	defer db.rwm.Unlock()
	_, ok := db.m[key]
	db.m[key] = val
	return !ok, nil
}
