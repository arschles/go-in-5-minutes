package main

import (
	"log"

	"github.com/arschles/go-in-5-minutes/episode9/conf"
)

func main() {
	// create the dedicated implementation
	mem := conf.NewMem()
	// use the default implementation (the singleton)
	conf.SetInt("a", 1)
	// use the dedicated implementation
	mem.SetInt("a", 2)

	// the remaining lines prove that the dedicated implementation is separate from the default (singleton) implementation
	i, err := conf.GetInt("a")
	if err != nil {
		log.Printf("(default) error getting int 'a' [%s]", err)
	} else {
		log.Printf("(default) int 'a' = %d", i)
	}
	i, err = mem.GetInt("a")
	if err != nil {
		log.Printf("(custom) error getting int 'a' [%s]", err)
	} else {
		log.Printf("(custom) int 'a' = %d", i)
	}
}
