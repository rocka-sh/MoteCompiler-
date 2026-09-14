package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// * (multiplicacion)

var LexemaMultiplicacion = adf.Lexema{
	QInicial: Q0Multiplicacion,
	Token:    "MULTIPLICACION",
}

var Q1Multiplicacion = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Multiplicacion = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'*': &Q1Multiplicacion,
	},
	IsF: false,
}
