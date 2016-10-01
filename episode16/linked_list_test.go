package main

import (
	"testing"
)

// TestBasic shows how to use subtests to set up the test environment, name tests properly, and then clean up
func TestBasic(t *testing.T) {
	ll := newLinkedList("a")
	t.Run("len", func(t *testing.T) {
		if ll.len() != 1 {
			t.Fatalf("length was not 1")
		}
	})
}
