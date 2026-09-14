package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// < (menor que)

var LexemaMenor = adf.Lexema{
	QInicial: Q0Menor,
	Token:    "MENOR",
}

var Q1Menor = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Menor = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'<': &Q1Menor,
	},
	IsF: false,
}
