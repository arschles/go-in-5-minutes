package main

type setElement struct {
	val string
}

// set is a simple set type. none of the functions on this type are concurrency safe
type set struct {
	m map[setElement]struct{}
}

// creates a new, empty set
func newSet() *set {
	return &set{m: make(map[setElement]struct{})}
}

// returns true if elt already exists in s
func (s *set) exists(elt setElement) bool {
	_, ok := s.m[elt]
	return ok
}

// adds elt to s unless it already exists. returns true if it was added successfully
func (s *set) add(elt setElement) bool {
	_, ok := s.m[elt]
	s.m[elt] = struct{}{}
	return !ok
}

// removes elt from s unless it doesn't exist. returns true if it was removed, false otherwise
func (s *set) remove(elt setElement) bool {
	_, ok := s.m[elt]
	if !ok {
		return false
	}
	delete(s.m, elt)
	return true
}

// removeAll removes all elements from s and returns the total number of elements removed
func (s *set) removeAll() int {
	ret := len(s.m)
	s.m = make(map[setElement]struct{})
	return ret
}
