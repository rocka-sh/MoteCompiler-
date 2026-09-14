package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaIn = adf.Lexema{
	QInicial: Q0In,
	Token:    "palabra_reservada",
}

var Q2In = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1In = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'n': &Q2In,
		'N': &Q2In,
	},
	IsF: false,
}

var Q0In = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'i': &Q1In,
		'I': &Q1In,
	},
	IsF: false,
}
