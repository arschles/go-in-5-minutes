package models

import (
	"encoding/json"
	"sync"
)

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
	b, ok := db.m[key]
	if !ok {
		return ErrNotFound
	}
	return json.Unmarshal(b, val)
}

func (db *inMemDB) Set(key string, val Model) error {
	db.rwm.Lock()
	defer db.rwm.Unlock()
	_, ok := db.m[key]
	if !ok {
		return ErrNotFound
	}
	b, err := val.MarshalJSON()
	if err != nil {
		return err
	}
	db.m[key] = b
	return nil
}

func (db *inMemDB) Upsert(key string, val Model) (bool, error) {
	db.rwm.Lock()
	defer db.rwm.Unlock()
	_, ok := db.m[key]
	b, err := val.MarshalJSON()
	if err != nil {
		return false, err
	}
	db.m[key] = b
	return !ok, nil
}
