package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// ( (parentesis abre)

var LexemaParentesisAbre = adf.Lexema{
	QInicial: Q0ParentesisAbre,
	Token:    "PARENTESIS_ABRE",
}

var Q1ParentesisAbre = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0ParentesisAbre = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'(': &Q1ParentesisAbre,
	},
	IsF: false,
}
