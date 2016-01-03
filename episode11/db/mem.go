package db

import (
	"errors"
	"sync"

	"github.com/arschles/go-in-5-minutes/episode11/models"
)

var ErrNotFound = errors.New("not found")

type Mem struct {
	mut sync.RWMutex
	m   map[string]models.Model
}

func NewMem() *Mem {
	return &Mem{m: make(map[string]models.Model)}
}

func (m *Mem) Save(key models.Key, model models.Model) error {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.m[key.String()] = model
	return nil
}

func (m *Mem) Delete(key models.Key) error {
	m.mut.Lock()
	defer m.mut.Unlock()
	delete(m.m, key.String())
	return nil
}

func (m *Mem) Get(key models.Key, model models.Model) error {
	m.mut.RLock()
	defer m.mut.RUnlock()
	md, ok := m.m[key.String()]
	if !ok {
		return ErrNotFound
	}
	model = md
	return nil
}
