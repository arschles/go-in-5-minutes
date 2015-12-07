package main

import (
	"log"

	"github.com/arschles/go-in-5-minutes/episode9/conf"
)

func main() {
	mem := conf.NewMem()
	conf.SetInt("a", 1)
	mem.SetInt("a", 2)
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
