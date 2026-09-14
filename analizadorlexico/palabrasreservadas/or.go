package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaOr = adf.Lexema{
	QInicial: Q0Or,
	Token:    "palabra_reservada",
}

var Q2Or = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1Or = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'r': &Q2Or,
		'R': &Q2Or,
	},
	IsF: false,
}

var Q0Or = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'o': &Q1Or,
		'O': &Q1Or,
	},
	IsF: false,
}
