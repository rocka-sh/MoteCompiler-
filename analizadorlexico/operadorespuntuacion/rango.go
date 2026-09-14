package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// .. (rango)

var LexemaRango = adf.Lexema{
	QInicial: Q0Rango,
	Token:    "rango",
}

var Q2Rango = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1Rango = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'.': &Q2Rango,
	},
	IsF: false,
}

var Q0Rango = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'.': &Q1Rango,
	},
	IsF: false,
}
