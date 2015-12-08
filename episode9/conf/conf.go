// Package conf is a configuration package. It can be used to store configuration data in multiple different pluggable backends
package conf

import (
	"errors"
	"sync"
)

// Default is an in-memory config backend. It's used in the SetInt and GetInt calls, and it's the optional singleton. Some packages export this name and others don't. If it's exported, it's the library user's responsibility to set it correctly and avoid concurrency issues (like setting it from two different goroutines, which would be a race condition)
var Default Data = NewMem()

// ErrNotFound is the error returned when a config item isn't found
var ErrNotFound = errors.New("not found")

// SetInt calls Default.SetInt
func SetInt(name string, i int) {
	Default.SetInt(name, i)
}

// GetInt calls Default.GetInt
func GetInt(name string) (int, error) {
	return Default.GetInt(name)
}

// Data is the core config interface
type Data interface {
	// SetInt sets the config value at name to i, overwriting if it already exists
	SetInt(name string, i int)
	// GetInt gets the config value at name. Returns 0, ErrNotFound if not such value found
	GetInt(name string) (int, error)
}

type memData struct {
	l       *sync.RWMutex
	strings map[string]string
	ints    map[string]int
}

// NewMem creates a Data implementation that stores config data in memory
func NewMem() Data {
	return &memData{l: &sync.RWMutex{}, strings: make(map[string]string), ints: make(map[string]int)}
}

// SetInt is the interface implementation
func (m *memData) SetInt(name string, i int) {
	m.l.Lock()
	defer m.l.Unlock()
	m.ints[name] = i
}

// GetInt is the interface implementation
func (m *memData) GetInt(name string) (int, error) {
	m.l.RLock()
	defer m.l.RUnlock()
	i, ok := m.ints[name]
	if !ok {
		return 0, ErrNotFound
	}
	return i, nil
}
