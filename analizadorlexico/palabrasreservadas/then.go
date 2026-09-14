package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaThen = adf.Lexema{
	QInicial: Q0Then,
	Token:    "palabra_reservada",
}

var Q4Then = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q3Then = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'n': &Q4Then,
		'N': &Q4Then,
	},
	IsF: false,
}

var Q2Then = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q3Then,
		'E': &Q3Then,
	},
	IsF: false,
}

var Q1Then = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'h': &Q2Then,
		'H': &Q2Then,
	},
	IsF: false,
}

var Q0Then = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		't': &Q1Then,
		'T': &Q1Then,
	},
	IsF: false,
}
