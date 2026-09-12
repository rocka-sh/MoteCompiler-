package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaElse = adf.Lexema{
	QInicial: Q0Else,
	Token:    "Palabra reservada",
}

var Q4Else = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q3Else = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q4Else,
		'E': &Q4Else,
	},
	IsF: false,
}

var Q2Else = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		's': &Q3Else,
		'S': &Q3Else,
	},
	IsF: false,
}

var Q1Else = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'l': &Q2Else,
		'L': &Q2Else,
	},
	IsF: false,
}

var Q0Else = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q1Else,
		'E': &Q1Else,
	},
	IsF: false,
}
