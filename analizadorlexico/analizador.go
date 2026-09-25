package analizadorlexico

import (
	"fmt"
	"os"
	"unicode"
)

func EscanearTokens(contenido []rune, tR Token, tL Token, tI Token, tO Token) {
	i := 0
	n := len(contenido)
	linea := 1

	for i < n {
		if contenido[i] == '\n' {
			linea++
			i++
			continue
		}
		if unicode.IsSpace(contenido[i]) {
			i++
			continue
		}

		if esComentario, avance := ignorarComentario(contenido, i, &linea); esComentario {
			i += avance
			continue
		}

		validarErroresPrevios(contenido, i, linea, tL)

		subslice := contenido[i:]
		var mejorLexema *Lexema
		mejorLongitud := 0

		if lex, l := tR.EvaluarPrefijo(subslice); l > mejorLongitud {
			mejorLongitud, mejorLexema = l, lex
		}
		if lex, l := tO.EvaluarPrefijo(subslice); l > mejorLongitud {
			mejorLongitud, mejorLexema = l, lex
		}
		if lex, l := tL.EvaluarPrefijo(subslice); l > mejorLongitud {
			mejorLongitud, mejorLexema = l, lex
		}
		if lex, l := tI.EvaluarPrefijo(subslice); l > mejorLongitud {
			mejorLongitud, mejorLexema = l, lex
		}

		if mejorLongitud > 0 {
			fmt.Println("Token:", mejorLexema.Token,
				"Lexema:", string(contenido[i:i+mejorLongitud]))
			validarErroresNumericos(contenido, i, mejorLongitud, mejorLexema, linea)
			i += mejorLongitud
		} else {
			fmt.Printf("Error lexico: caracter no reconocido: '%c' en linea %d\n", contenido[i], linea)
			os.Exit(1)
			i++
		}
	}
}
