package identificadores

import adf "MoteCompiler/analizadorlexico"

var Identificadores = adf.Tokenario{
	Tipo: "Identificador",
	Tokens: []*adf.Token{
		&LexemaIdentificador,
	},
}

func InitADF() {
	var r rune

	Q1 := &adf.Estado{
		IsF: true,
	}

	mapaId0 := make(map[rune]*adf.Estado)
	mapaId1 := make(map[rune]*adf.Estado)

	for r = 'a'; r <= 'z'; r++ {
		mapaId0[r] = Q1
	}
	for r = 'A'; r <= 'Z'; r++ {
		mapaId0[r] = Q1
	}
	mapaId0['_'] = Q1

	for r = 'a'; r <= 'z'; r++ {
		mapaId1[r] = Q1
	}
	for r = 'A'; r <= 'Z'; r++ {
		mapaId1[r] = Q1
	}
	mapaId1['_'] = Q1
	for r = '0'; r <= '9'; r++ {
		mapaId1[r] = Q1
	}

	Q1.Transiciones = mapaId1

	Q0 := adf.Estado{
		Transiciones: mapaId0,
		IsF:          false,
	}

	LexemaIdentificador = adf.Token{
		QInicial:    Q0,
		NombreToken: "token_identificador",
	}
}
