package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaEnd = adf.Lexema{
	QInicial: Q0End,
	Token:    "Palabra reservada",
}

var Q3End = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q2End = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'd': &Q3End,
		'D': &Q3End,
	},
	IsF: false,
}

var Q1End = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'n': &Q2End,
		'N': &Q2End,
	},
	IsF: false,
}

var Q0End = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q1End,
		'E': &Q1End,
	},
	IsF: false,
}
