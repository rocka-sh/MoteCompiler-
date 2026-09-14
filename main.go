package main

import (
	adf "MoteCompiler/analizadorlexico"
	id "MoteCompiler/analizadorlexico/identificadores"
	lit "MoteCompiler/analizadorlexico/literales"
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

	EscanearTokens(palabras, pr.PalabrasReservadas, lit.Literales, id.Identificadores)

}

func EscanearTokens(contenido []rune, tokenReservadas adf.Token, tokenLiterales adf.Token, tokenIdentificador adf.Token) {
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

		// 1. Palabras Reservadas
		lexRes, lenRes := tokenReservadas.EvaluarPrefijo(subslice)
		if lenRes > mejorLongitud {
			mejorLongitud = lenRes
			mejorLexema = lexRes
		}

		// 2. Literales
		lexLit, lenLit := tokenLiterales.EvaluarPrefijo(subslice)
		if lenLit > mejorLongitud {
			mejorLongitud = lenLit
			mejorLexema = lexLit
		}

		// 3. Identificador (solo si supera en longitud a la palabra reservada)
		lexId, lenId := tokenIdentificador.EvaluarPrefijo(subslice)
		if lenId > mejorLongitud {
			mejorLongitud = lenId
			mejorLexema = lexId
		}

		if mejorLongitud > 0 {
			fmt.Println("Token:", mejorLexema.Token, "Lexema:", string(contenido[i:i+mejorLongitud]))
			i += mejorLongitud
		} else {
			fmt.Println("Caracteres no reconocidos:", string(contenido[i]))
			i++
		}
	}
}
