package main

import (
	"log"

	"github.com/arschles/go-in-5-minutes/episode17/pkg"
)

func checkAndPrint(i, j int) {
	res, err := pkg.Div(i, j)
	if err != nil {
		if pkg.IsDivByZeroErr(err) {
			log.Printf("divide by zero error (%s)", err)
		} else if pkg.IsDivByOneErr(err) {
			log.Printf("divide by one error (%s)", err)
		} else if pkg.IsDivisorGTDividendErr(err) {
			log.Printf("divisor > dividend error (%s)", err)
		}
		return
	}
	log.Printf("successful divide: %d / %d = %d", i, j, res)
}

func main() {
	// should have a divide by 0 error
	checkAndPrint(1, 0)
	// should have a divide by 1 error
	checkAndPrint(2, 1)
	// should have a divisor > dividend error
	checkAndPrint(1, 2)
	// should succeed
	checkAndPrint(4, 2)
}
