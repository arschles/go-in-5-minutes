package main

import (
	"math/rand"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type dish struct {
	name       string
	numMorsels int
}

var (
	dishes = []dish{
		dish{name: "chorizo", numMorsels: 5 + rand.Intn(5)},
		dish{name: "chopitos", numMorsels: 5 + rand.Intn(5)},
		dish{name: "pimientos de padrón", numMorsels: 5 + rand.Intn(5)},
		dish{name: "croquetas", numMorsels: 5 + rand.Intn(5)},
		dish{name: "patatas bravas", numMorsels: 5 + rand.Intn(5)},
	}
)

func randomDish() dish {
	return dishes[rand.Intn(len(dishes))]
}
