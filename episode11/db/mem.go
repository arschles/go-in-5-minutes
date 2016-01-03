package db

import (
	"errors"
	"sync"

	"github.com/arschles/go-in-5-minutes/episode11/models"
)

// ErrNotFound is returned when a key is passed that is not in the in-memory database
var ErrNotFound = errors.New("not found")

// Mem is an in-memory only implementation of DB
type Mem struct {
	mut sync.RWMutex
	m   map[string]models.Model
}

// NewMem initializes and returns an empty Mem database
func NewMem() *Mem {
	return &Mem{m: make(map[string]models.Model)}
}

// Save is the interface implementation. Overwrites existing keys and never returns an error
func (m *Mem) Save(key models.Key, model models.Model) error {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.m[key.String()] = model
	return nil
}

// Delete is the interface implementation. Never returns an error, even if the key didn't exist
func (m *Mem) Delete(key models.Key) error {
	m.mut.Lock()
	defer m.mut.Unlock()
	delete(m.m, key.String())
	return nil
}

// Get is the interface implementation. Returns ErrNotFound if no such key existed. Callers should pass a pointer to a model so that Get can write the fetched model into it
func (m *Mem) Get(key models.Key, model models.Model) error {
	m.mut.RLock()
	defer m.mut.RUnlock()
	md, ok := m.m[key.String()]
	if !ok {
		return ErrNotFound
	}
	return model.Set(md)
}
