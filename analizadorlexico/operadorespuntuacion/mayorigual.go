package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// >= (mayor o igual)

var LexemaMayorIgual = adf.Lexema{
	QInicial: Q0MayorIgual,
	Token:    "mayor_igual",
}

var Q2MayorIgual = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1MayorIgual = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'=': &Q2MayorIgual,
	},
	IsF: false,
}

var Q0MayorIgual = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'>': &Q1MayorIgual,
	},
	IsF: false,
}
