package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// = (igual)

var LexemaIgual = adf.Lexema{
	QInicial: Q0Igual,
	Token:    "IGUAL",
}

var Q1Igual = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Igual = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'=': &Q1Igual,
	},
	IsF: false,
}
