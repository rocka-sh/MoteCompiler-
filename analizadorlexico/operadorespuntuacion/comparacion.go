package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// <> (comparacion / diferente)

var LexemaComparacion = adf.Lexema{
	QInicial: Q0Comparacion,
	Token:    "comparacion",
}

var Q2Comparacion = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1Comparacion = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'>': &Q2Comparacion,
	},
	IsF: false,
}

var Q0Comparacion = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'<': &Q1Comparacion,
	},
	IsF: false,
}
