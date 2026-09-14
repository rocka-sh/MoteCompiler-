package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// [ (corchete abre)

var LexemaCorcheteAbre = adf.Lexema{
	QInicial: Q0CorcheteAbre,
	Token:    "corchete_abre",
}

var Q1CorcheteAbre = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0CorcheteAbre = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'[': &Q1CorcheteAbre,
	},
	IsF: false,
}
