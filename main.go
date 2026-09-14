package main

import (
	adf "MoteCompiler/analizadorlexico"
	id "MoteCompiler/analizadorlexico/identificadores"
	lit "MoteCompiler/analizadorlexico/literales"
	op "MoteCompiler/analizadorlexico/operadorespuntuacion"
	pr "MoteCompiler/analizadorlexico/palabrasreservadas"
	"fmt"
	"os"
	"unicode"
)

func main() {
	argumentos := os.Args
	palabras, err := adf.LectorArchivo(argumentos[1])
	if err != nil {
		fmt.Println("No se pudo abrir el archivo:", err)
		return
	}

	tR := pr.PalabrasReservadas
	tL := lit.Literales
	tI := id.Identificadores
	tO := op.OperadoresPuntuacion

	lit.InitADF()
	id.InitADF()

	EscanearTokens(palabras, tR, tL, tI, tO)

}

func EscanearTokens(contenido []rune, tR adf.Token, tL adf.Token, tI adf.Token, tO adf.Token) {
	i := 0
	n := len(contenido)

	for i < n {
		if unicode.IsSpace(contenido[i]) {
			i++
			continue
		}

		subslice := contenido[i:]
		var mejorLexema *adf.Lexema
		mejorLongitud := 0

		lexRes, lenRes := tR.EvaluarPrefijo(subslice)
		if lenRes > mejorLongitud {
			mejorLongitud = lenRes
			mejorLexema = lexRes
		}

		lexOp, lenOp := tO.EvaluarPrefijo(subslice)
		if lenOp > mejorLongitud {
			mejorLongitud = lenOp
			mejorLexema = lexOp
		}

		lexLit, lenLit := tL.EvaluarPrefijo(subslice)
		if lenLit > mejorLongitud {
			mejorLongitud = lenLit
			mejorLexema = lexLit
		}

		lexId, lenId := tI.EvaluarPrefijo(subslice)
		if lenId > mejorLongitud {
			mejorLongitud = lenId
			mejorLexema = lexId
		}

		if mejorLongitud > 0 {
			fmt.Println("Token:", mejorLexema.Token,
				"Lexema:", string(contenido[i:i+mejorLongitud]))
			i += mejorLongitud
		} else {
			fmt.Println("Caracteres no reconocidos:", string(contenido[i]))
			i++
		}
	}
}
