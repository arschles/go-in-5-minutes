package models

import (
	"fmt"
)

// Key is the generic key that's used to identify Models in db.DB implementations
type Key interface {
	fmt.Stringer
}
