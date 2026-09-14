package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaString = adf.Lexema{
	QInicial: Q0String,
	Token:    "Palabra reservada",
}

var Q6String = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q5String = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'g': &Q6String,
		'G': &Q6String,
	},
	IsF: false,
}

var Q4String = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'n': &Q5String,
		'N': &Q5String,
	},
	IsF: false,
}

var Q3String = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'i': &Q4String,
		'I': &Q4String,
	},
	IsF: false,
}

var Q2String = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'r': &Q3String,
		'R': &Q3String,
	},
	IsF: false,
}

var Q1String = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		't': &Q2String,
		'T': &Q2String,
	},
	IsF: false,
}

var Q0String = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		's': &Q1String,
		'S': &Q1String,
	},
	IsF: false,
}
