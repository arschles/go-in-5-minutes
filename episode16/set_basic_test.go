package main

import (
	"fmt"
	"strconv"
	"testing"
)

// TestBasic shows how to use subtests to set up the test environment & name tests properly
func TestBasic(t *testing.T) {
	const num = 10
	st := newSet()
	for i := 0; i < num; i++ {
		elt := setElement{val: strconv.Itoa(i)}
		st.add(elt)
	}
	defer st.removeAll()
	for i := 0; i < num; i++ {
		expectedElt := setElement{val: strconv.Itoa(i)}
		testName := fmt.Sprintf("element%d", i)
		t.Run(testName, func(t *testing.T) {
			if !st.exists(expectedElt) {
				t.Errorf("set reported %s as missing", expectedElt.val)
			}
		})
	}
}
