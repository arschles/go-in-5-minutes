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
	// runs once before the entire test suite runs. you could so something in this function
	// like set up some global state. here are a few common tasks to consider:
	//
	// - initialize the global random number generator:
	//		rand.Seed(time.Now().UnixNano())
	// - create & seed a database
}

func (t *TransformerTests) SetupTest() {
	// runs once before every individual test runs. you can do things like clear common
	// state so that everything is "clean" for the test to run. be careful if you're
	// doing things to an external database, though. by default, tests run in parallel
	// so if all your tests share the same database and tables, they'll be overwriting
	// each other.
	//
	// you can fix that by either using unique DB/table names or running the tests without
	// parallelism (go test -p 1 ./...)
	t.transformer = &transformer{
		mut:   new(sync.Mutex),
		cache: map[string]interface{}{},
	}
}

func (t *TransformerTests) TearDownTest() {
	// runs once after every individual test runs. you can clean up after your tests here.
	// a common thing here is to shut down an HTTP server that you left running
}

func (t *TransformerTests) TearDownSuite() {
	// runs once after the entire test suite runs. a common thing to do here is shut down
	// a database you created before the tests ran
}

func TestTransformer(t *testing.T) {
	suite.Run(t, &TransformerTests{
		Suite: new(suite.Suite),
	})
}
