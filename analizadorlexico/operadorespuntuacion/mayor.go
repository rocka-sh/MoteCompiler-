package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// > (mayor que)

var LexemaMayor = adf.Lexema{
	QInicial: Q0Mayor,
	Token:    "token_mayor",
}

var Q1Mayor = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Mayor = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'>': &Q1Mayor,
	},
	IsF: false,
}
