package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// % (modulo)

var LexemaModulo = adf.Lexema{
	QInicial: Q0Modulo,
	Token:    "modulo",
}

var Q1Modulo = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Modulo = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'%': &Q1Modulo,
	},
	IsF: false,
}
