package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaElif = adf.Token{
	QInicial:    Q0Elif,
	NombreToken: "token_elif",
}

var Q4Elif = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q3Elif = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'f': &Q4Elif,
		'F': &Q4Elif,
	},
	IsF: false,
}

var Q2Elif = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'i': &Q3Elif,
		'I': &Q3Elif,
	},
	IsF: false,
}

var Q1Elif = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'l': &Q2Elif,
		'L': &Q2Elif,
	},
	IsF: false,
}

var Q0Elif = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q1Elif,
		'E': &Q1Elif,
	},
	IsF: false,
}
