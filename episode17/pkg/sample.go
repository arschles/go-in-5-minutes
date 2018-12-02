package pkg

import (
	"errors"
	"fmt"
)

var (
	errDivZero = errors.New("divisor is zero")
	errDivOne  = errors.New("divisor is one")
)

type errDivisorGTDividend struct {
	dividend int
	divisor  int
}

func (e errDivisorGTDividend) Error() string {
	return fmt.Sprintf("divisor (%d) is greater than dividend (%d)", e.divisor, e.dividend)
}

// IsDivByZeroErr returns true if e is a divide by zero error
func IsDivByZeroErr(e error) bool {
	return e == errDivZero
}

// IsDivByOneErr returns true if e is a divide by one error
func IsDivByOneErr(e error) bool {
	return e == errDivOne
}

// IsDivisorGTDividendErr returns true if e is a divisor > dividend error
func IsDivisorGTDividendErr(e error) bool {
	_, ok := e.(errDivisorGTDividend)
	return ok
}

// Div returns dividend/divisor. Returns 0 and one of the following errors in the following cases:
//
//  - If divisor is zero, returns an error that you can check with IsDivByZeroErr
//  - If divisor is one, returns an error that you can check with IsDivByOneErr
//  - If divisor > dividend, returns an error that you can check with IsDivisorGTDividendErr
func Div(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errDivZero
	} else if divisor == 1 {
		return 0, errDivOne
	} else if divisor > dividend {
		return 0, errDivisorGTDividend{dividend: dividend, divisor: divisor}
	}
	return dividend / divisor, nil
}
