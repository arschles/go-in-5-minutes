package models

import (
	"encoding"
)

type Model interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
