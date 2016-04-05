package etl

import "strings"

// Transform cambio de formato de los datos antiguos.
// in  -> int = array de letras.
// out -> char = int
// asigna a cada letra del array dado la puntación asignada.
func Transform(m map[int][]string) map[string]int {
	solucion := make(map[string]int)
	for puntos, letras := range m {
		for _, letra := range letras {
			solucion[strings.ToLower(letra)] = puntos
		}
	}
	return solucion
}
