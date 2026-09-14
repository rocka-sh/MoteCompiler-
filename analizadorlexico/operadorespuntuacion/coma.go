package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// , (coma)

var LexemaComa = adf.Lexema{
	QInicial: Q0Coma,
	Token:    "COMA",
}

var Q1Coma = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Coma = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		',': &Q1Coma,
	},
	IsF: false,
}
