package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaTrue = adf.Lexema{
	QInicial: Q0True,
	Token:    "Palabra reservada",
}

var Q4True = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q3True = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q4True,
		'E': &Q4True,
	},
	IsF: false,
}

var Q2True = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'u': &Q3True,
		'U': &Q3True,
	},
	IsF: false,
}

var Q1True = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'r': &Q2True,
		'R': &Q2True,
	},
	IsF: false,
}

var Q0True = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		't': &Q1True,
		'T': &Q1True,
	},
	IsF: false,
}
