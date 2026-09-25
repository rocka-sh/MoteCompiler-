package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

// <= (menor o igual)

var LexemaMenorIgual = adf.Token{
	QInicial:    Q0MenorIgual,
	NombreToken: "token_menor_igual",
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
