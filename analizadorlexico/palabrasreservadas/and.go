package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaAnd = adf.Lexema{
	QInicial: Q0And,
	Token:    "palabra_reservada",
}

var Q3And = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q2And = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'd': &Q3And,
		'D': &Q3And,
	},
	IsF: false,
}

var Q1And = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'n': &Q2And,
		'N': &Q2And,
	},
	IsF: false,
}

var Q0And = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'a': &Q1And,
		'A': &Q1And,
	},
	IsF: false,
}
