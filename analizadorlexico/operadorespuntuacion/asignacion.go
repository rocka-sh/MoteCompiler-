package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// := (asignacion)

var LexemaAsignacion = adf.Token{
	QInicial:    Q0Asignacion,
	NombreToken: "token_asignacion",
}

var Q2Asignacion = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1Asignacion = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'=': &Q2Asignacion,
	},
	IsF: false,
}

var Q0Asignacion = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		':': &Q1Asignacion,
	},
	IsF: false,
}
