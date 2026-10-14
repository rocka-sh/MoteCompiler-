package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// . (punto / acceso a campos)

var LexemaPunto = adf.Token{
	QInicial:    Q0Punto,
	NombreToken: ".",
}

var Q1Punto = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Punto = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'.': &Q1Punto,
	},
	IsF: false,
}
