package episode0

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

// HashTable is the interface for a simple hash table. It's designed in such a way
// that some libraries immediately adhere to it with no extra code.
// One such library is https://godoc.org/github.com/hoisie/redis
type HashTable interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
