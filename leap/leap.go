// Leap stub file

// This next comment is a "build constraint."  It's here for, well, kind of
// complicated testing purposes you don't need to worry about now.
// Actually, you can delete it.

// +build !example

// The package name is expected by the test program.
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// It's good style to write a comment here documenting IsLeapYear.
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
