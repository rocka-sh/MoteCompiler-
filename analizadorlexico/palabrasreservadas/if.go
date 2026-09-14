package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaIf = adf.Lexema{
	QInicial: Q0If,
	Token:    "palabra_reservada",
}

var Q2If = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1If = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'f': &Q2If,
		'F': &Q2If,
	},
	IsF: false,
}

var Q0If = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'i': &Q1If,
		'I': &Q1If,
	},
	IsF: false,
}
