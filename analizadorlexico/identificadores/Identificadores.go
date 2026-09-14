package identificadores

import adf "MoteCompiler/analizadorlexico"

var Q1Identificador = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Identificador = adf.Estado{
	Transiciones: nil,
	IsF:          false,
}

var LexemaIdentificador = adf.Lexema{
	QInicial: Q0Identificador,
	Token:    "IDENTIFICADOR",
}
