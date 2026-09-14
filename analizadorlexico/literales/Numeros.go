package literales

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaEntero = adf.Lexema{
	QInicial: Q0Entero,
	Token:    "Numero Entero",
}

var LexemaFlotante = adf.Lexema{
	QInicial: Q0Flotante,
	Token:    "Numero Flotante",
}

var Q1Entero = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q0Entero = adf.Estado{
	Transiciones: nil,
	IsF:          false,
}

var Q3Flotante = adf.Estado{
	Transiciones: nil,
	IsF:          true,
}

var Q2Flotante = adf.Estado{
	Transiciones: nil,
	IsF:          false,
}

var Q1Flotante = adf.Estado{
	Transiciones: map[rune]*adf.Estado{
		'.': &Q2Flotante,
	},
	IsF: false,
}

var Q0Flotante = adf.Estado{
	Transiciones: nil,
	IsF:          false,
}

// GO NO PERMITE INICIALIZAR STRUCTS QUE SE LLAMEN A ELLOS MISMOS, POR ESO SE HACE EN EL INIT
// PRINCIPALEMENTE PARA LOS NUMEROS, CADENAS E IDENTIFICADORES, YA QUE PUEDEN SER INFINITOS

func initNumeros() {
	// INICIALIZACION DE NUMEROS PARA ENTEROS
	var mapaN1 map[rune]*adf.Estado = make(map[rune]*adf.Estado)
	var mapaN0 map[rune]*adf.Estado = make(map[rune]*adf.Estado)
	for r := '0'; r <= '9'; r++ {
		mapaN0[r] = &Q1Entero
		mapaN1[r] = &Q1Entero
	}
	Q0Entero.Transiciones = mapaN0
	Q1Entero.Transiciones = mapaN1

	// INICIALIZACION DE NUMEROS PARA FLOTANTES
	var mapaF2 map[rune]*adf.Estado = make(map[rune]*adf.Estado)
	var mapaF3 map[rune]*adf.Estado = make(map[rune]*adf.Estado)
	var mapaF0 map[rune]*adf.Estado = make(map[rune]*adf.Estado)

	for r := '0'; r <= '9'; r++ {
		mapaF0[r] = &Q1Flotante
		mapaF2[r] = &Q3Flotante
		mapaF3[r] = &Q3Flotante

	}
	Q0Flotante.Transiciones = mapaF0
	Q2Flotante.Transiciones = mapaF2
	Q3Flotante.Transiciones = mapaF3

}
