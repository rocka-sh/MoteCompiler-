package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaDo = adf.Token{
	QInicial:    Q0Do,
	NombreToken: "token_do",
}

var Q2Do = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1Do = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'o': &Q2Do,
		'O': &Q2Do,
	},
	IsF: false,
}

var Q0Do = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'd': &Q1Do,
		'D': &Q1Do,
	},
	IsF: false,
}
