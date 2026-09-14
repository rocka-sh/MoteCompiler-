package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// - (resta)

var LexemaResta = adf.Lexema{
	QInicial: Q0Resta,
	Token:    "RESTA",
}

var Q1Resta = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Resta = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'-': &Q1Resta,
	},
	IsF: false,
}
