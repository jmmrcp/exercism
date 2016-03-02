// Package secret permite convertir un numero entero
// en una serie de strings o nil sino se cumplen las
// condiciones.
package secret

import "strconv"

// ConvtoBin convierte el entero a reprentacion binaria en texto.
//
// convtoBin int -> string
// ConvtoBin(19)
// "10011"
func convtoBin(n int) string {
	f := int64(n)
	return strconv.FormatInt(f, 2)
}

// Stringtobool convierte representacion binaria a slice booleano.
//
// stringtobool string -> []bool
// Stringtobool("10011")
//[true,false,false,true,true]
func stringtobool(s string) []bool {
	var t = []bool{false, false, false, false}
	var y bool
	for i := 0; i < len(s); i++ {
		if s[i] == 49 {
			y = true
		} else {
			y = false
		}
		t = append(t, y)
	}
	if len(t) > 5 {
		dif := len(t) - 5
		cp := t[dif:]
		return cp
	}
	return t
}

// Handshake devuelve un slice de texto con el codigo del entero
//
// Handshake int -> []string
// Handshake(19)
// ["double blink", "wink"]
func Handshake(n int) []string {
	var code []string
	stringBinario := convtoBin(n)
	if n < 1 {
		return code
	}
	stringBooleano := stringtobool(stringBinario)
	x := stringBooleano
	if x[4] == true {
		code = append(code, "wink")
	}
	if x[3] == true {
		code = append(code, "double blink")
	}
	if x[2] == true {
		code = append(code, "close your eyes")
	}
	if x[1] == true {
		code = append(code, "jump")
	}
	if x[0] == true {
		if len(code) > 0 {
			code = invstring(code)
		}
	}
	return code
}

// invstring ...
func invstring(s []string) []string {
	var sliceInvertida []string
	for i := len(s); i > 0; i-- {
		sliceInvertida = append(sliceInvertida, s[i-1])
	}
	return sliceInvertida
}
