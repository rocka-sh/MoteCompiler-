package identificadores

import adf "MoteCompiler/analizadorlexico"

var Identificadores = adf.Token{
	Tipo: "Identificador",
	Lexemas: []*adf.Lexema{
		&LexemaIdentificador,
	},
}

func InitADF() {
	var r rune
	var mapaId0 map[rune]*adf.Estado = make(map[rune]*adf.Estado)
	var mapaId1 map[rune]*adf.Estado = make(map[rune]*adf.Estado)

	for r = 'a'; r <= 'z'; r++ {
		mapaId0[r] = &Q1Identificador
	}
	for r = 'A'; r <= 'Z'; r++ {
		mapaId0[r] = &Q1Identificador
	}
	mapaId0['_'] = &Q1Identificador

	for r = 'a'; r <= 'z'; r++ {
		mapaId1[r] = &Q1Identificador
	}
	for r = 'A'; r <= 'Z'; r++ {
		mapaId1[r] = &Q1Identificador
	}
	mapaId1['_'] = &Q1Identificador
	for r = '0'; r <= '9'; r++ {
		mapaId1[r] = &Q1Identificador
	}

	Q0Identificador.Transiciones = mapaId0
	Q1Identificador.Transiciones = mapaId1
}
