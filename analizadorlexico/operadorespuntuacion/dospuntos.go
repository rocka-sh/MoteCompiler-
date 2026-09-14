package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// : (dos puntos)

var LexemaDosPuntos = adf.Lexema{
	QInicial: Q0DosPuntos,
	Token:    "DOS_PUNTOS",
}

var Q1DosPuntos = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0DosPuntos = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		':': &Q1DosPuntos,
	},
	IsF: false,
}
