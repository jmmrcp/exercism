// +build !example

package raindrops

import "strconv"

const testVersion = 2

func Convert(i int) string {
	var str string
	if i%3 == 0 {
		str += "Pling"
	}
	if i%5 == 0 {
		str += "Plang"
	}
	if i%7 == 0 {
		str += "Plong"
	}
	if str == "" {
		str += strconv.Itoa(i)
	}
	return str
}

// The test program has a benchmark too.  How fast does your Convert convert?
