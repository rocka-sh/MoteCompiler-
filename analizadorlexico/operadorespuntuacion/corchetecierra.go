package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// ] (corchete cierra)

var LexemaCorcheteCierra = adf.Lexema{
	QInicial: Q0CorcheteCierra,
	Token:    "corchete_cierra",
}

var Q1CorcheteCierra = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0CorcheteCierra = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		']': &Q1CorcheteCierra,
	},
	IsF: false,
}
