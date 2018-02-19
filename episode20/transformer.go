package episode20

import (
	"sync"
)

type transformer struct {
	mut   *sync.Mutex
	cache map[string]interface{}
}

// looks up key in the cache. if it exists, runs its value through fn, puts the new value
// back into the cache, and returns true. otherwise does nothing and returns false
func (t *transformer) transform(key string, fn func(string) string) bool {
	t.mut.Lock()
	defer t.mut.Unlock()
	val, found := t.cache[key]
	if !found {
		return false
	}
	str, ok := val.(string)
	if !ok {
		return false
	}
	newStr := fn(str)
	t.cache[key] = newStr
	return true
}

// removes key from the cache if it existed and returns true. otherwise returns false
func (t *transformer) remove(key string) bool {
	t.mut.Lock()
	defer t.mut.Unlock()
	if _, found := t.cache[key]; found {
		delete(t.cache, key)
		return true
	}
	return false
}
