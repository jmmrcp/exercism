package grains

import "fmt"

// Square ...
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("%d", n)
	}
	return 1 << uint64(n-1), nil
}

// Total ...
func Total() uint64 {
	return (1 << 64) - 1
}
