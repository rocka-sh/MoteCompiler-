package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// ) (parentesis cierra)

var LexemaParentesisCierra = adf.Token{
	QInicial:    Q0ParentesisCierra,
	NombreToken: "token_parentesis_cierra",
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
