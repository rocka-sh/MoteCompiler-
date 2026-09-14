package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaPrint = adf.Lexema{
	QInicial: Q0Print,
	Token:    "palabra_reservada",
}

var Q5Print = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q4Print = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		't': &Q5Print,
		'T': &Q5Print,
	},
	IsF: false,
}

var Q3Print = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'n': &Q4Print,
		'N': &Q4Print,
	},
	IsF: false,
}

var Q2Print = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'i': &Q3Print,
		'I': &Q3Print,
	},
	IsF: false,
}

var Q1Print = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'r': &Q2Print,
		'R': &Q2Print,
	},
	IsF: false,
}

var Q0Print = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'p': &Q1Print,
		'P': &Q1Print,
	},
	IsF: false,
}
