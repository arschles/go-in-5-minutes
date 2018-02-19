package episode20

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TransformerTests struct {
	*suite.Suite
	transformer *transformer
}

func (t *TransformerTests) SetupSuite() {
	// logic in here will run once, before all tests run
	t.transformer = &transformer{
		mut:   new(sync.Mutex),
		cache: map[string]interface{}{},
	}
}

func (t *TransformerTests) SetupTest() {
	// logic in here will run before every individual test runs
	t.transformer = &transformer{
		mut:   new(sync.Mutex),
		cache: map[string]interface{}{},
	}
}

func (t *TransformerTests) TearDownTest() {
	// logic in here will run after every individual test runs
}

func (t *TransformerTests) TearDownSuite() {
	// logic in here will run once, after all tests run
}

func TestTransformer(t *testing.T) {
	suite.Run(t, &TransformerTests{
		Suite: new(suite.Suite),
	})
}
