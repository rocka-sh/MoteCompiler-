package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// < (menor que)

var LexemaMenor = adf.Token{
	QInicial:    Q0Menor,
	NombreToken: "token_menor",
}

var Q1Menor = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Menor = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'<': &Q1Menor,
	},
	IsF: false,
}
