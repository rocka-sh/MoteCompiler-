package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// ( (parentesis abre)

var LexemaParentesisAbre = adf.Token{
	QInicial:    Q0ParentesisAbre,
	NombreToken: "token_parentesis_abre",
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
