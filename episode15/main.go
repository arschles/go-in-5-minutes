package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numComputers      = 8
	numTourists       = 25
	minOnlineDuration = 5 * time.Second
	maxOnlineDuration = 10 * time.Second
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func logf(fmtStr string, params ...interface{}) {
	fmt.Printf(fmtStr+"\n", params...)
}

func pluralize(i int, singular, plural string) string {
	if i == 1 {
		return singular
	}
	return plural
}

func sleepRand() time.Duration {
	randSec := maxOnlineDuration.Seconds() - minOnlineDuration.Seconds()
	sleepSec := rand.Intn(int(randSec)) + int(minOnlineDuration.Seconds())
	time.Sleep(time.Duration(sleepSec) * time.Second)
	return time.Duration(sleepSec) * time.Second
}

// this struct represents a tourist who is going to use a computer
type tourist struct {
	// the tourist number
	num int
}

func main() {
	// this channel represents the line of people waiting outside the doors. each value passed through the channel is the tourist number
	queue := make(chan tourist)
	// this channel is closed when all of the people are done using the computers. we're using this to broadcast to all the computer goroutines to tell them to stop
	stopCh := make(chan struct{})

	// create the computers. each computer is represented by a goroutine
	for i := 0; i < numComputers; i++ {
		go func(computerNum int) {
			for {
				select {
				case <-stopCh:
				case tst := <-queue:
					logf("Tourist %d is online.", tst.num)
					spent := sleepRand()
					mins := int(spent.Seconds())
					logf(
						"Tourist %d is done, having spent %d %s online.",
						tst.num,
						mins,
						pluralize(mins, "second", "seconds"),
					)
				}
			}
		}(i)
	}

	// Open the doors and start sending people on the queue channel. We're assuming there is no line, and the first person who gets to the internet cafe before doors open will not necessarily be the first who gets a computer
	for i := 1; i <= numTourists; i++ {
		tst := tourist{num: i}
		go func(touristNum int) {
			// we're using select here to either send immediately, or if we can't, log and then try to send
			select {
			case queue <- tst:
			default:
				logf("Tourist %d is waiting for turn.", touristNum)
				queue <- tst
			}
		}(i)
	}
	logf("The place is empty, let's close up and go to the beach!")
	close(stopCh)
}
