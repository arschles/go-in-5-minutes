package dishes

import "math/rand"

// Manager is an actor that manages the internal list of dishes and number of morsels left in each.
type Manager struct {
	// The slice of shared dishes. This is the shared state that the Manager actor is managing.
	dishes []dish
	// This is the message passing channel.
	//
	// The manager receive loop only reads from it (in a 'range Ch') and writes to the 'chan *string' channel values.
	//
	// The message sender must only write to Ch and then read from the channel values. Simply close this channel to terminate the receive loop.
	//
	// The manager will write a valid dish name back on the 'chan *string' if there are any left, or nil if there are none left.
	Ch chan chan *string
}

// NewManager creates a new Manager actor and starts the receive loop in a background goroutine.
//
// Example usage:
//
//  mgr := NewManager()
//  // get a dish that has morsels left
//  dishNameCh := make(chan *string)
//  mgr.Ch <- dishNameCh
//  dishName := <-dishNameCh
//  if dishName == nil {
//    // there are no more morsels left of any dish!
//  } else {
//    fmt.Println("got a morsel of %s!", *dishName)
//  }
func NewManager() *Manager {
	ch := make(chan chan *string)
	mgr := &Manager{
		// this is the shared state that we're managing
		dishes: []dish{
			{name: "chorizo", numMorsels: randNumMorsels()},
			{name: "chopitos", numMorsels: randNumMorsels()},
			{name: "pimientos de padrón", numMorsels: randNumMorsels()},
			{name: "croquetas", numMorsels: randNumMorsels()},
			{name: "patatas bravas", numMorsels: randNumMorsels()},
		},
		Ch: ch,
	}

	go func() {
		for retCh := range mgr.Ch {
			if len(mgr.dishes) == 0 {
				retCh <- nil
				continue
			}
			idx := rand.Intn(len(mgr.dishes))
			dish := &mgr.dishes[idx]
			dish.numMorsels--
			if dish.numMorsels == 0 {
				// remove the dish from the list
				mgr.dishes = append(mgr.dishes[0:idx], mgr.dishes[idx+1:]...)
			}
			retCh <- &dish.name
		}
	}()
	return mgr
}

// randDishIdx returns a random index into the dishes slice, or -1 if the slice
// is empty
func (m *Manager) randDishIdx() int {
	if len(m.dishes) == 0 {
		return -1
	}
	return rand.Intn(len(m.dishes))
}

type dish struct {
	name       string
	numMorsels int
}

// randNumMorsels returns a random number of morsels between 5 and 10
func randNumMorsels() int {
	return 5 + rand.Intn(5)
}
