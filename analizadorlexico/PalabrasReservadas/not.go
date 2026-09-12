package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaNot = adf.Lexema{
	QInicial: Q0Not,
	Token:    "Palabra reservada",
}

var Q3Not = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q2Not = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		't': &Q3Not,
		'T': &Q3Not,
	},
	IsF: false,
}

var Q1Not = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'o': &Q2Not,
		'O': &Q2Not,
	},
	IsF: false,
}

var Q0Not = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'n': &Q1Not,
		'N': &Q1Not,
	},
	IsF: false,
}
