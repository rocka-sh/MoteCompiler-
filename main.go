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

		if validarErroresPrevios(contenido, i, linea, tL) {
			i += consumirHastaEspacio(contenido, i)
			continue
		}

		subslice := contenido[i:]
		var mejorLexema *adf.Lexema
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
			if validarErroresNumericos(contenido, i, mejorLongitud, mejorLexema, linea) {
				i += consumirHastaEspacio(contenido, i)
				continue
			}

			fmt.Println("Token:", mejorLexema.Token, "Lexema:", string(contenido[i:i+mejorLongitud]))
			i += mejorLongitud
		} else {
			fmt.Printf("Error lexico: caracter no reconocido: '%c' en línea %d\n", contenido[i], linea)
			i++
		}
	}
}

//funciones auxiliares para los errores lexicos y el escaneo de tokens

func esLetraOSimboloId(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

// Procesa #... y #|...|#
func ignorarComentario(contenido []rune, i int, linea *int) (bool, int) {
	n := len(contenido)
	if contenido[i] != '#' {
		return false, 0
	}

	// Comentario de Bloque
	if i+1 < n && contenido[i+1] == '|' {
		pos := i + 2
		cerrado := false
		for pos < n {
			if contenido[pos] == '\n' {
				*linea++
			}
			if contenido[pos] == '|' && pos+1 < n && contenido[pos+1] == '#' {
				pos += 2
				cerrado = true
				break
			}
			pos++
		}
		if !cerrado {
			fmt.Printf("Error lexico: comentario de bloque sin cerrar linea %d\n", *linea)
		}
		return true, pos - i
	}

	// Comentario de Línea
	pos := i
	for pos < n && contenido[pos] != '\n' {
		pos++
	}
	return true, pos - i
}

// "sin cerrar y .X
func validarErroresPrevios(contenido []rune, i int, linea int, tL adf.Token) bool {
	n := len(contenido)

	if contenido[i] == '"' {
		lexLit, lenLit := tL.EvaluarPrefijo(contenido[i:])
		if lenLit == 0 || lexLit.Token != "LITERAL_CADENA" {
			fmt.Printf("Error lexico: cadena sin cerrar en linea %d\n", linea)
			return true
		}
	}

	if contenido[i] == '.' && i+1 < n && contenido[i+1] >= '0' && contenido[i+1] <= '9' {
		fmt.Printf("Error lexico: literal flotante mal formado en linea %d\n", linea)
		return true
	}

	return false
}

// Revisa 3abc o 3.
func validarErroresNumericos(contenido []rune, i int, mejorLongitud int, mejorLexema *adf.Lexema, linea int) bool {
	if mejorLexema.Token != "numero_entero" && mejorLexema.Token != "numero_flotante" {
		return false
	}

	n := len(contenido)
	posSig := i + mejorLongitud

	if posSig < n && esLetraOSimboloId(contenido[posSig]) {
		fmt.Printf("Error lexico: identificador no puede comenzar con digito en linea %d\n", linea)
		return true
	}

	if posSig < n && contenido[posSig] == '.' {
		if posSig+1 >= n || contenido[posSig+1] != '.' {
			fmt.Printf("Error lexico: literal flotante mal formado en linea %d\n", linea)
			return true
		}
	}

	return false
}

func consumirHastaEspacio(contenido []rune, i int) int {
	pos := i
	for pos < len(contenido) && !unicode.IsSpace(contenido[pos]) {
		pos++
	}
	return pos - i
}
