package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// / (division)

var LexemaDivision = adf.Lexema{
	QInicial: Q0Division,
	Token:    "division",
}

var Q1Division = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Division = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'/': &Q1Division,
	},
	IsF: false,
}
