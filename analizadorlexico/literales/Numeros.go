package literales

import (
	adf "MoteCompiler/analizadorlexico"
)

var LexemaEntero adf.Lexema
var LexemaFlotante adf.Lexema

// GO NO PERMITE INICIALIZAR STRUCTS QUE SE LLAMEN A ELLOS MISMOS, POR ESO SE HACE EN EL INIT
// PRINCIPALEMENTE PARA LOS NUMEROS, CADENAS E IDENTIFICADORES, YA QUE PUEDEN SER INFINITOS

func initNumeros() {
	// INICIALIZACION DE NUMEROS PARA ENTEROS
	Q1E := &adf.Estado{
		IsF: true,
	}

	mapaN0 := make(map[rune]*adf.Estado)
	mapaN1 := make(map[rune]*adf.Estado)
	for r := '0'; r <= '9'; r++ {
		mapaN0[r] = Q1E
		mapaN1[r] = Q1E
	}
	Q1E.Transiciones = mapaN1

	Q0E := adf.Estado{
		Transiciones: mapaN0,
		IsF:          false,
	}

	LexemaEntero = adf.Lexema{
		QInicial: Q0E,
		Token:    "Numero Entero",
	}

	// INICIALIZACION DE NUMEROS PARA FLOTANTES
	Q3F := &adf.Estado{
		IsF: true,
	}

	Q2F := &adf.Estado{
		IsF: false,
	}

	Q1F := &adf.Estado{
		Transiciones: map[rune]*adf.Estado{
			'.': Q2F,
		},
		IsF: false,
	}

	mapaF0 := make(map[rune]*adf.Estado)
	mapaF2 := make(map[rune]*adf.Estado)
	mapaF3 := make(map[rune]*adf.Estado)

	for r := '0'; r <= '9'; r++ {
		mapaF0[r] = Q1F
		mapaF2[r] = Q3F
		mapaF3[r] = Q3F
	}

	// Q1F también necesita las transiciones de dígitos para números como 123.1
	for r := '0'; r <= '9'; r++ {
		Q1F.Transiciones[r] = Q1F
	}

	Q2F.Transiciones = mapaF2
	Q3F.Transiciones = mapaF3

	Q0F := adf.Estado{
		Transiciones: mapaF0,
		IsF:          false,
	}

	LexemaFlotante = adf.Lexema{
		QInicial: Q0F,
		Token:    "Numero Flotante",
	}
}

