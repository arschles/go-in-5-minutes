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

// seed the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// convenience function for calling fmt.Printf with a trailing newline
func logf(fmtStr string, params ...interface{}) {
	fmt.Printf(fmtStr+"\n", params...)
}

// convenience function that returns plural if i != 1, singular otherwise
func pluralize(i int, singular, plural string) string {
	if i == 1 {
		return singular
	}
	return plural
}

// convenience function to sleep for a random amount of time in [minOnlineDuration, maxOnlineDuration]
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
				// when we close the stop channel (at the end of the main function), every one of the computer goroutines will receive here, and then return from this function. effectively, that means that the computer will "shut down" and stop waiting for tourists to come
				case <-stopCh:
					return
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
	// HACK: we're sleeping an arbitrary amount of time here to let tourists finish using the computer. The proper fix is in the extended screencast. We can increase the time to let more tourists finish, but the code won't always wait for all tourists (since they take a random amount of time to finish using the computer)
	//
	// The extended screencast shows how to fix the code so that it _deterministically_ exits when all tourists are done. See https://gum.co/gifm-x-15 for more
	time.Sleep(1 * time.Second)
	logf("The place is empty, let's close up and go to the beach!")
	close(stopCh)
}
