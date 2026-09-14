package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaRecord = adf.Lexema{
	QInicial: Q0Record,
	Token:    "palabra_reservada",
}

var Q6Record = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q5Record = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'd': &Q6Record,
		'D': &Q6Record,
	},
	IsF: false,
}

var Q4Record = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'r': &Q5Record,
		'R': &Q5Record,
	},
	IsF: false,
}

var Q3Record = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'o': &Q4Record,
		'O': &Q4Record,
	},
	IsF: false,
}

var Q2Record = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'c': &Q3Record,
		'C': &Q3Record,
	},
	IsF: false,
}

var Q1Record = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q2Record,
		'E': &Q2Record,
	},
	IsF: false,
}

var Q0Record = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'r': &Q1Record,
		'R': &Q1Record,
	},
	IsF: false,
}
