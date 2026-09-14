package literales

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaCadena adf.Lexema

func initCadenas() {
	Q2 := &adf.Estado{
		Transiciones: nil,
		IsF:          true,
	}

	Q1 := &adf.Estado{
		IsF: false,
	}

	mapaQ1 := make(map[rune]*adf.Estado)
	for r := rune(32); r <= 126; r++ {
		if r != '"' && r != '\\' {
			mapaQ1[r] = Q1
		}
	}
	mapaQ1['"'] = Q2
	Q1.Transiciones = mapaQ1

	Q0 := adf.Estado{
		Transiciones: map[rune]*adf.Estado{
			'"': Q1,
		},
		IsF: false,
	}

	LexemaCadena = adf.Lexema{
		QInicial: Q0,
		Token:    "literal_cadena",
	}
}
