package models

import "errors"

var ErrNotFound = errors.New("not found")

// DB is the interface to the database. The idea here is to have multiple
// different DB implementations. For example, you could build an in-memory
// one for testing and a MongoDB one for production
type DB interface {
	// Get gets the value for the given key. returns the value and nil
	// if it was found, nil and ErrNotFound if it was not found and another
	// appropriate error otherwise
	Get(key string) ([]byte, error)
	// Set sets the value for the given existing key. returns nil if the key
	// already existed and was successfully set, ErrNotFound if the key didn't
	// already exist, and another appropriate error otherwise
	Set(key string, val []byte) error
	// Upsert sets or creates the value for the existing key. returns true and nil
	// if the key was created on this call, false and nil if the key was not created
	// but still successfully updated, and false and the appropriate error otherwise
	Upsert(key string, val []byte) (bool, error)
}
