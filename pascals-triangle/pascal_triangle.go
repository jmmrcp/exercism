package pascal

// Triangle ...
func Triangle(n int) [][]int {

	valor := make([][]int, n)
	for i := 0; i < n; i++ {
		valor[i] = make([]int, i+1)
	}
	for i := 0; i < n; i++ {
		valor[i][0] = 1
	}
	for i := 0; i < n; i++ {
		valor[i][i] = 1
	}
	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			valor[i][j] = valor[i-1][j] + valor[i-1][j-1]
		}
	}
	return valor
}
