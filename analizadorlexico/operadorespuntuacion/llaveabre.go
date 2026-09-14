package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// { (llave abre)

var LexemaLlaveAbre = adf.Lexema{
	QInicial: Q0LlaveAbre,
	Token:    "llave_abre",
}

var Q1LlaveAbre = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0LlaveAbre = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'{': &Q1LlaveAbre,
	},
	IsF: false,
}
