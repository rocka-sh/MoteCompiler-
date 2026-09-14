package literales

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaFalse = adf.Lexema{
	QInicial: Q0False,
	Token:    "booleano",
}

var Q5False = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q4False = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'e': &Q5False,
		'E': &Q5False,
	},
	IsF: false,
}

var Q3False = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		's': &Q4False,
		'S': &Q4False,
	},
	IsF: false,
}

var Q2False = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'l': &Q3False,
		'L': &Q3False,
	},
	IsF: false,
}

var Q1False = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'a': &Q2False,
		'A': &Q2False,
	},
	IsF: false,
}

var Q0False = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'f': &Q1False,
		'F': &Q1False,
	},
	IsF: false,
}
