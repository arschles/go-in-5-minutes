package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("starting 'heavy duty' work")
	go recur(1, []string{"1"})

	log.Println("serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	os.Exit(1)
}

// heavyDuty represents some "heavy duty" work to be done. this will grow memory usage
// and goroutine usage forever because it does recursion
func recur(i int, strs []string) {
	time.Sleep(30 * time.Second)
	s1 := append(strs, strconv.Itoa(i))
	s2 := append(s1, strconv.Itoa(i+1))
	go recur(i+1, s1)
	recur(i+2, s2)
}
