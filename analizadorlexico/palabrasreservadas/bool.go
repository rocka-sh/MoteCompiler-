package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaBool = adf.Lexema{
	QInicial: Q0Bool,
	Token:    "Palabra reservada",
}

var Q4Bool = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q3Bool = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'l': &Q4Bool,
		'L': &Q4Bool,
	},
	IsF: false,
}

var Q2Bool = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'o': &Q3Bool,
		'O': &Q3Bool,
	},
	IsF: false,
}

var Q1Bool = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'o': &Q2Bool,
		'O': &Q2Bool,
	},
	IsF: false,
}

var Q0Bool = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'b': &Q1Bool,
		'B': &Q1Bool,
	},
	IsF: false,
}
