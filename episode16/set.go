package main

type setElement struct {
	value string
}

type set struct {
	m map[setElement]struct{}
}

func newSet() *set {
	return &set{m: make(map[setElement]struct{})}
}

func (s *set) exists(elt setElement) bool {
	_, ok := s.m[elt]
	return ok
}

func (s *set) add(elt setElement) {
	s.m[elt] = struct{}{}
}
