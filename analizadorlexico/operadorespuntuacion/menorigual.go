package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// <= (menor o igual)

var LexemaMenorIgual = adf.Lexema{
	QInicial: Q0MenorIgual,
	Token:    "MENOR_IGUAL",
}

var Q2MenorIgual = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q1MenorIgual = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'=': &Q2MenorIgual,
	},
	IsF: false,
}

var Q0MenorIgual = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'<': &Q1MenorIgual,
	},
	IsF: false,
}
