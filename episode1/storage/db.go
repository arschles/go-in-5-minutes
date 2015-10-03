package storage

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type DB interface {
	// Get returns the value for the given key, ErrNotFound if the key doesn't exist,
	// or another error if the get failed
	Get(key string) ([]byte, error)
	// Set sets the value for the given key. Returns an error if the set failed.
	// If non-nil error is returned, the value was not updated
	Set(key string, val []byte) error
}
