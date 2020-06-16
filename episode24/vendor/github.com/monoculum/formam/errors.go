package formam

type Error struct {
	err error
}

func (s *Error) Error() string {
	return "formam: " + s.err.Error()
}

func newError(err error) *Error { return &Error{err} }
