package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// = (igual)

var LexemaIgual = adf.Token{
	QInicial:    Q0Igual,
	NombreToken: "token_igual",
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
