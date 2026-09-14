package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// } (llave cierra)

var LexemaLlaveCierra = adf.Lexema{
	QInicial: Q0LlaveCierra,
	Token:    "llave_cierra",
}

var Q1LlaveCierra = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0LlaveCierra = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'}': &Q1LlaveCierra,
	},
	IsF: false,
}
