package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/arschles/go-in-5-minutes/episode12/dishes"
)

func init() {
	// set the number of OS threads to use to the number of CPUs to use
	runtime.GOMAXPROCS(runtime.NumCPU())
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

// randEatTime returns a random time between 30 seconds and 3 minutes (180 seconds)
func randEatTime(minSec, maxSec int) time.Duration {
	sec := rand.Intn(maxSec-minSec) + minSec
	return time.Duration(sec) * time.Second
}

func main() {
	minEatSec := flag.Int("min", 30, "the minimum number of seconds it takes someone to eat a morsel")
	maxEatSec := flag.Int("max", 180, "the maximum number of seconds it takes someone to eat a morsel")
	flag.Parse()
	if *minEatSec > *maxEatSec {
		fmt.Printf("Error: min (%d) is greater than max (%d)\n", *minEatSec, *maxEatSec)
		os.Exit(1)
	}
	names := []string{"Alice", "Bob", "Charlie", "Dave"}
	fmt.Println("Bon appétit!")
	mgr := dishes.NewManager()
	var wg sync.WaitGroup
	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			for {
				dishNameCh := make(chan *string)
				mgr.Ch <- dishNameCh
				dishName := <-dishNameCh
				if dishName == nil {
					return
				}
				fmt.Printf("%s is enjoying some %s\n", name, *dishName)
				time.Sleep(randEatTime(*minEatSec, *maxEatSec))
			}
		}(name)
	}

	wg.Wait()
	// this shuts down the manager
	close(mgr.Ch)
	fmt.Println("That was delicious!")
}
