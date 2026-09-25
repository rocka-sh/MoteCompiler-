package main

import (
	adf "MoteCompiler/analizadorlexico"
	id "MoteCompiler/analizadorlexico/identificadores"
	lit "MoteCompiler/analizadorlexico/literales"
	op "MoteCompiler/analizadorlexico/operadorespuntuacion"
	pr "MoteCompiler/analizadorlexico/palabrasreservadas"
)

func Lexico(contenido []rune) []adf.Token {
	tR := pr.PalabrasReservadas
	tL := lit.Literales
	tI := id.Identificadores
	tO := op.OperadoresPuntuacion

	lit.InitADF()
	id.InitADF()

	return adf.EscanearTokens(contenido, tR, tL, tI, tO)

}
