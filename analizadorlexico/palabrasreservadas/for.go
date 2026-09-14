package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaFor = adf.Lexema{
	QInicial: Q0For,
	Token:    "Palabra reservada",
}

var Q3For = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q2For = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'r': &Q3For,
		'R': &Q3For,
	},
	IsF: false,
}

var Q1For = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'o': &Q2For,
		'O': &Q2For,
	},
	IsF: false,
}

var Q0For = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'f': &Q1For,
		'F': &Q1For,
	},
	IsF: false,
}
