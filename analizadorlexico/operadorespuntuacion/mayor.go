package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// > (mayor que)

var LexemaMayor = adf.Token{
	QInicial:    Q0Mayor,
	NombreToken: ">",
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
