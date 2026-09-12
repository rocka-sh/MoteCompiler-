package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaVar = adf.Lexema{
	QInicial: Q0Var,
	Token:    "Palabra reservada",
}

var Q3Var = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q2Var = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'r': &Q3Var,
		'R': &Q3Var,
	},
	IsF: false,
}

var Q1Var = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'a': &Q2Var,
		'A': &Q2Var,
	},
	IsF: false,
}

var Q0Var = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'v': &Q1Var,
		'V': &Q1Var,
	},
	IsF: false,
}
