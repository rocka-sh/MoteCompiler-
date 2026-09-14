package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaReturn = adf.Lexema{
	QInicial: Q0Return,
	Token:    "palabra_reservada",
}

var Q6Return = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q5Return = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'n': &Q6Return,
		'N': &Q6Return,
	},
	IsF: false,
}

var Q4Return = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'r': &Q5Return,
		'R': &Q5Return,
	},
	IsF: false,
}

var Q3Return = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'u': &Q4Return,
		'U': &Q4Return,
	},
	IsF: false,
}

var Q2Return = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		't': &Q3Return,
		'T': &Q3Return,
	},
	IsF: false,
}

var Q1Return = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q2Return,
		'E': &Q2Return,
	},
	IsF: false,
}

var Q0Return = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'r': &Q1Return,
		'R': &Q1Return,
	},
	IsF: false,
}
