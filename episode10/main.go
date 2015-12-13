package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	errInvalidRange = errors.New("invalid range")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// getReady tells name to get ready. when they're done, it calls wg.Done() to indicate to anyone listening that they're done. this func also prints messages to indicate that name has started and finished getting ready.
func getReady(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s started getting ready\n", name)
	// the time it takes to get ready is at least 60 seconds, plus a random number up to 90 seconds
	sec := rand.Intn(30) + 60
	// this is where they get ready
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Printf("%s spent %d seconds getting ready\n", name, sec)
}

// putOnShoes tells name to put on their shoes. when they're done, it calls wg.Done() to indicate to anyone listening that they're done. this func also prints messages to indicate that name has started and finished putting on their shoes.
func putOnShoes(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s started putting on shoes\n", name)
	// the time it takes to put on shoes is at least 35 seconds, plus a random number up to 45 seconds
	sec := rand.Intn(10) + 35
	// this is where they put on their shoes
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Printf("%s spent %d seconds putting on shoes\n", name, sec)
}

// armAlarm starts the arming process and notifies when it's finished.
// This func does the following in order:
//
//  - waits for startArm to receive
//  - closes armStarted to indicate to everyone listening that the arming process has started
//  - sleeps for 60 seconds to do the delay
//  - when the delay has finished, closes armFinished to signal that the alarm is armed
//
// Run this func in a goroutine and interact with it using the channels that it's passed.
// See an example in the main below.
func armAlarm(startArm <-chan struct{}, armStarted chan<- struct{}, armFinished chan<- struct{}) {
	<-startArm        // wait for the signal to start arming
	close(armStarted) // indicate that the signal has been received and the arming countdown has started
	time.Sleep(60 * time.Second)
	close(armFinished)
}

func main() {
	// set the number of OS threads to use to the number of CPUs to use
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Let's go for a walk!")
	var wg sync.WaitGroup
	wg.Add(2)
	// tell Alice and Bob to get ready
	go getReady("Alice", &wg)
	go getReady("Bob", &wg)
	// wait for Alice and Bob to finish getting ready
	wg.Wait()

	// arm the alarm
	startArm := make(chan struct{})
	armStarted := make(chan struct{})
	armFinished := make(chan struct{})
	go armAlarm(startArm, armStarted, armFinished)
	close(startArm) // tell the alarm to start the arming process
	fmt.Println("Arming alarm.")
	<-armStarted // wait for the alarm to have started arming

	// the WaitGroup is 0 after the above Wait() call returns. reset it to 3 to get notified that:
	// - bot of them have finished putting on their shoes
	// - the "Alarm is counting down." message has finished printing
	wg.Add(3)
	go func() {
		// notify that the alarm is counting down
		fmt.Println("Alarm is counting down.")
		wg.Done()
	}()

	// have Alice and Bob put on their shoes
	go putOnShoes("Alice", &wg)
	go putOnShoes("Bob", &wg)

	wg.Wait() // wait for Alice and Bob to finish putting on their shoes
	fmt.Println("Exiting and locking the door.")

	<-armFinished // wait for the alarm to finish arming
	fmt.Println("Alarm is Armed.")
}
