package models

import "sync"

type inMemDB struct {
	rwm *sync.RWMutex
	m   map[string]Model
}

func NewInMemoryDB() DB {
	return &inMemDB{
		rwm: &sync.RWMutex{},
		m:   make(map[string]Model),
	}
}

func (db *inMemDB) GetAllKeys() ([]string, error) {
	db.rwm.RLock()
	defer db.rwm.RUnlock()
	ret := make([]string, len(db.m))
	i := 0
	for key, _ := range db.m {
		ret[i] = key
		i++
	}
	return ret, nil
}

func (db *inMemDB) Get(key string, val Model) error {
	db.rwm.RLock()
	defer db.rwm.RUnlock()
	v, ok := db.m[key]
	if !ok {
		return ErrNotFound
	}
	val = v
	return nil
}

func (db *inMemDB) Set(key string, val Model) error {
	db.rwm.Lock()
	defer db.rwm.Unlock()
	_, ok := db.m[key]
	if !ok {
		return ErrNotFound
	}
	db.m[key] = val
	return nil
}

func (db *inMemDB) Upsert(key string, val Model) (bool, error) {
	db.rwm.Lock()
	defer db.rwm.Unlock()
	_, ok := db.m[key]
	db.m[key] = val
	return !ok, nil
}
