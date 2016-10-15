package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestTableDriven(t *testing.T) {
	elts := []setElement{
		setElement{val: strconv.Itoa(rand.Int())},
		setElement{val: strconv.Itoa(rand.Int())},
		setElement{val: strconv.Itoa(rand.Int())},
		setElement{val: strconv.Itoa(rand.Int())},
		setElement{val: strconv.Itoa(rand.Int())},
		setElement{val: strconv.Itoa(rand.Int())},
	}
	st := newSet()
	for i, elt := range elts {
		t.Run(fmt.Sprintf("element%d", i), func(t *testing.T) {
			if st.exists(elt) {
				t.Errorf("element %d reported as existing", i)
				return
			}
			if !st.add(elt) {
				t.Errorf("element %d reported as existing", i)
			}
			if !st.exists(elt) {
				t.Errorf("element %d reported as missing", i)
				return
			}
			if !st.remove(elt) {
				t.Errorf("element %d reported as not removed", i)
				return
			}
		})
	}
}
