package analizadorlexico

import (
	"fmt"
	"os"
	"unicode"
)

func EscanearTokens(contenido []rune, tR Tokenario, tL Tokenario, tI Tokenario, tO Tokenario) []Token {
	var tokens []Token
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
		var mejorToken *Token
		mejorLongitud := 0

		if lex, l := tR.EvaluarPrefijo(subslice); l > mejorLongitud {
			mejorLongitud, mejorToken = l, lex
		}
		if lex, l := tO.EvaluarPrefijo(subslice); l > mejorLongitud {
			mejorLongitud, mejorToken = l, lex
		}
		if lex, l := tL.EvaluarPrefijo(subslice); l > mejorLongitud {
			mejorLongitud, mejorToken = l, lex
		}
		if lex, l := tI.EvaluarPrefijo(subslice); l > mejorLongitud {
			mejorLongitud, mejorToken = l, lex
		}

		if mejorLongitud > 0 {
			validarErroresNumericos(contenido, i, mejorLongitud, mejorToken, linea)
			tok := Token{
				QInicial:    mejorToken.QInicial,
				NombreToken: mejorToken.NombreToken,
				Lexema:      string(contenido[i : i+mejorLongitud]),
				Linea:       linea,
			}

			tokens = append(tokens, tok)
			i += mejorLongitud
		} else {
			fmt.Printf("Error lexico: caracter no reconocido: '%c' en linea %d\n", contenido[i], linea)
			os.Exit(1)
			i++
		}
	}

	return tokens
}
