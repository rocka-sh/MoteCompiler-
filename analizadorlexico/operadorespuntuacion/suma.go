package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// + (suma)

var LexemaSuma = adf.Lexema{
	QInicial: Q0Suma,
	Token:    "SUMA",
}

var Q1Suma = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Suma = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'+': &Q1Suma,
	},
	IsF: false,
}
