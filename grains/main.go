package main

import "fmt"

func Square(n int) (uint64, error) {
	if n < 1 {
		return 0, fmt.Errorf("El numero debe ser mayor de %d.", n)
	}
	return 1 << uint64(n-1), nil
}

func Total() uint64 {
	return (1 << 64) - 1
}

func main() {
	fmt.Printf("%d.\n", Total())
}
