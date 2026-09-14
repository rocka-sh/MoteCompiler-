package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// ) (parentesis cierra)

var LexemaParentesisCierra = adf.Lexema{
	QInicial: Q0ParentesisCierra,
	Token:    "parentesis_cierra",
}

var Q1ParentesisCierra = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0ParentesisCierra = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		')': &Q1ParentesisCierra,
	},
	IsF: false,
}
