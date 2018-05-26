package validators

import (
	"fmt"
	"unicode/utf8"

	"github.com/gobuffalo/validate"
)

type StringLengthInRange struct {
	Name    string
	Field   string
	Min     int
	Max     int
	Message string
}

// IsValid checks that string in range of min:max
// if max not present or it equal to 0 it will be equal to string length
func (v *StringLengthInRange) IsValid(errors *validate.Errors) {
	strLength := utf8.RuneCountInString(v.Field)
	if v.Max == 0 {
		v.Max = strLength
	}
	if v.Message == "" {
		v.Message = fmt.Sprintf("%s not in range(%d, %d)", v.Name, v.Min, v.Max)
	}
	if !(strLength >= v.Min && strLength <= v.Max) {
		errors.Add(GenerateKey(v.Name), v.Message)
	}
}
