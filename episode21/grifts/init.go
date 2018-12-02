package grifts

import (
  "github.com/gobuffalo/buffalo"
	"github.com/arschles/go-in-5-minutes/episode21/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
