package models

import (
	"encoding"
)

type PrimaryKey string

func (p PrimaryKey) String() string {
	return string(p)
}

type Model interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	Key() PrimaryKey
}
