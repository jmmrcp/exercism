// Package diffsquares paquete matematico que calcula:
// El cuadrado de la suma de n.
// La suma de los cuadrados de n.
// La diferencia entre ambos.
package diffsquares

// SquareOfSums devuelve el Cuadrado de la suma de n numeros.
//
// SquareOfSums int -> int
// SquareofSums(10)
// 3025
func SquareOfSums(n int) int {
	var s1, s2, i int = 0, 0, 1
	for i <= n {
		s1 += i
		s2 = s1 * s1
		i++
	}
	return s2

}

// SumOfSquares devuelve la suma de los cuadrados de n numeros.
//
// SumOfSquares int -> int
// SumOfSquares(10)
// 385
func SumOfSquares(n int) int {
	s := 0
	i := 1
	for i <= n {
		s += i * i
		i++
	}
	return s

}

// Difference devulve la diferencia entre el cuadrado de la suma
// y la suma de los cuadrados de n numeros
//
// Difference uint -> uint
// Difference(10)
// 2640
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
