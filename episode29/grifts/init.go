package grifts

import (
	"github.com/arschles/go-in-5-minutes/episode29/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
