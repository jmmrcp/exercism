package slice

// All ...
func All(n int, s string) []string {

	if n < 1 || n > len(s) {
		var valor []string
		return valor
	}
	if n == 1 {
		l := len(s)
		valor := make([]string, l)
		for i := 0; i < l; i++ {
			valor[i] = string(s[i])
		}
		return valor
	}
	if n == 2 {
		l := len(s) - 1
		valor := make([]string, l)
		for i := 0; i < l; i++ {
			valor[i] = string(s[i]) + string(s[i+1])
		}
		return valor
	}
	if n == 3 {
		l := len(s) - 2
		valor := make([]string, l)
		for i := 0; i < l; i++ {
			valor[i] = string(s[i]) + string(s[i+1]) + string(s[i+2])
		}
		return valor

	}
	if n == 4 {
		l := len(s) - 3
		valor := make([]string, l)
		for i := 0; i < l; i++ {
			valor[i] = string(s[i]) + string(s[i+1]) + string(s[i+2]) + string(s[i+3])
		}
		return valor

	}
	if n == 5 {
		l := len(s) - 4
		valor := make([]string, l)
		for i := 0; i < l; i++ {
			valor[i] = string(s[i]) + string(s[i+1]) + string(s[i+2]) + string(s[i+3]) + string(s[i+4])
		}

		return valor

	}
	var valor []string
	return valor
}

// Frist ...
func Frist(n int, s string) string {
	var t string
	for i := 0; i <= n-1; i++ {
		t += string(s[i])
	}
	return t
}

// First ...
func First(n int, s string) (string, bool) {
	if n < 1 || n > len(s) {
		return "", false
	}
	var t string
	for i := 0; i <= n-1; i++ {
		t += string(s[i])
	}
	return t, true
}
