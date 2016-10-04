package main

type setElement struct {
	val string
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

func (s *set) add(elt setElement) bool {
	_, ok := s.m[elt]
	s.m[elt] = struct{}{}
	return !ok
}

func (s *set) remove(elt setElement) bool {
	_, ok := s.m[elt]
	if ok {
		return false
	}
	delete(s.m, elt)
	return true
}

func (s *set) removeAll() int {
	ret := len(s.m)
	s.m = make(map[setElement]struct{})
	return ret
}
