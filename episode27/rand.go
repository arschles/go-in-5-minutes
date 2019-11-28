package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randStr(slc []string) string {
	idx := rand.Intn(len(slc))
	return slc[idx]
}
