package actions

import (
	"testing"

	"github.com/gobuffalo/suite"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	// action, err := suite.NewActionWithFixtures(App(), packr.New("Test_ActionSuite", "../fixtures"))
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// as := &ActionSuite{
	// 	Action: action,
	// }
	// suite.Run(t, as)
}