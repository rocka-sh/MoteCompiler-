package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaLet = adf.Lexema{
	QInicial: Q0Let,
	Token:    "Palabra reservada",
}

var Q3Let = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q2Let = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		't': &Q3Let,
		'T': &Q3Let,
	},
	IsF: false,
}

var Q1Let = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q2Let,
		'E': &Q2Let,
	},
	IsF: false,
}

var Q0Let = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'l': &Q1Let,
		'L': &Q1Let,
	},
	IsF: false,
}
