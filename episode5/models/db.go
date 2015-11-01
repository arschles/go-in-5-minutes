package models

// DB is the interface to the database. The idea here is to have multiple
// different DB implementations. For example, you could build an in-memory
// one for testing and a MongoDB one for production
type DB interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte) error
	Upsert(key string, val []byte) (bool, error)
}
