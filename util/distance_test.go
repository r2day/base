package util

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	a := "22.21,129.292"
	b := "22.23,129.295"
	var want float64 = 2.245129
	msg := DistancePosition(a, b)
	if want != msg {
		t.Fatalf(`DistancePosition(a, b) = %f, want match for %#f, nil`, msg, want)
	}
}
