// Clock stub file

// Build constraint.
// It lets our CI (Continuous Integration) testing test the test program.
// But it's served its purpose by now.  Feel free to delete.
// +build !example

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 3

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock int // Complete the type definition.  Pick a suitable data type.

func New(hour, minute int) Clock {
	var minutos Clock
	minutos = Clock(hour*60 + minute)
	minutos %= 1440
	if minutos < 0 {
		minutos += 1440
	}
	return minutos
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

func (c Clock) Add(minutes int) Clock {
	var minutos Clock
	minutos = c + Clock(minutes)
	minutos %= 1440
	if minutos < 0 {
		minutos += 1440
	}
	return minutos
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
