package episode0

// BusinessLogic is a sample function that does some business logic using h as the data storage system.
// We will test this function
func BusinessLogic(h HashTable) {
	h.Set("hello", []byte("world"))
}
