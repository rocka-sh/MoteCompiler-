package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaInt = adf.Lexema{
	QInicial: Q0Int,
	Token:    "Palabra reservada",
}

var Q3Int = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q2Int = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		't': &Q3Int,
		'T': &Q3Int,
	},
	IsF: false,
}

var Q1Int = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'n': &Q2Int,
		'N': &Q2Int,
	},
	IsF: false,
}

var Q0Int = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'i': &Q1Int,
		'I': &Q1Int,
	},
	IsF: false,
}
