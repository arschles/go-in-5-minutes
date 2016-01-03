package models

import (
	"encoding"
)

// Model is the generic interface to represent data that's stored in db.DB implementations
type Model interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
