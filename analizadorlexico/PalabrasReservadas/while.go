package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaWhile = adf.Lexema{
	QInicial: Q0While,
	Token:    "Palabra reservada",
}

var Q5While = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q4While = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q5While,
		'E': &Q5While,
	},
	IsF: false,
}

var Q3While = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'l': &Q4While,
		'L': &Q4While,
	},
	IsF: false,
}

var Q2While = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'i': &Q3While,
		'I': &Q3While,
	},
	IsF: false,
}

var Q1While = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'h': &Q2While,
		'H': &Q2While,
	},
	IsF: false,
}

var Q0While = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'w': &Q1While,
		'W': &Q1While,
	},
	IsF: false,
}
