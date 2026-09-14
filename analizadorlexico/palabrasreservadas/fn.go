package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaFn = adf.Lexema{
	QInicial: Q0Fn,
	Token:    "palabra_reservada",
}

var Q2Fn = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1Fn = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'n': &Q2Fn,
		'N': &Q2Fn,
	},
	IsF: false,
}

var Q0Fn = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'f': &Q1Fn,
		'F': &Q1Fn,
	},
	IsF: false,
}
