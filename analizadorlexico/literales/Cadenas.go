package literales

import (
	adf "MoteCompiler/analizadorlexico"
)

var Q2Cadena = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1Cadena = adf.Estado{
	Transiciones: nil,
	IsF:          false,
}

var Q0Cadena = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'"': &Q1Cadena,
	},
	IsF: false,
}

var LexemaCadena = adf.Lexema{
	Token: "LITERAL_CADENA",
}

func initCadenas() {
	var r rune
	var mapaC0 map[rune]*adf.Estado = make(map[rune]*adf.Estado)
	for r = 32; r <= 126; r++ {
		if r != '"' && r != '\\' {
			mapaC0[r] = &Q1Cadena
		}
	}
	mapaC0['"'] = &Q2Cadena
	Q1Cadena.Transiciones = mapaC0
}
