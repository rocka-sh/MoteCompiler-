package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaFloat = adf.Lexema{
	QInicial: Q0Float,
	Token:    "Palabra reservada",
}

var Q5Float = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q4Float = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		't': &Q5Float,
		'T': &Q5Float,
	},
	IsF: false,
}

var Q3Float = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'a': &Q4Float,
		'A': &Q4Float,
	},
	IsF: false,
}

var Q2Float = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'o': &Q3Float,
		'O': &Q3Float,
	},
	IsF: false,
}

var Q1Float = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'l': &Q2Float,
		'L': &Q2Float,
	},
	IsF: false,
}

var Q0Float = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'f': &Q1Float,
		'F': &Q1Float,
	},
	IsF: false,
}
