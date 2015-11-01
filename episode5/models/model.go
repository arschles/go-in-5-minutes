package models

import "encoding/json"

type Model interface {
	json.Marshaler
}
