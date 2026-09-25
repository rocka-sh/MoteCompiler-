package analizadorlexico

import (
	"fmt"
	"os"
)

func esLetraOSimboloId(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func ignorarComentario(contenido []rune, i int, linea *int) (bool, int) {
	n := len(contenido)
	if contenido[i] != '#' {
		return false, 0
	}

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
			os.Exit(1)
		}
		return true, pos - i
	}

	pos := i
	for pos < n && contenido[pos] != '\n' {
		pos++
	}
	return true, pos - i
}

func validarErroresPrevios(contenido []rune, i int, linea int, tL Tokenario) {
	n := len(contenido)

	if contenido[i] == '"' {
		lexLit, lenLit := tL.EvaluarPrefijo(contenido[i:])
		if lenLit == 0 || lexLit.NombreToken != "token_cadena" {
			fmt.Printf("Error lexico: cadena sin cerrar en linea %d\n", linea)
			os.Exit(1)
		}
	}

	if contenido[i] == '.' && i+1 < n && contenido[i+1] >= '0' && contenido[i+1] <= '9' {
		fmt.Printf("Error lexico: literal flotante mal formado en linea %d\n", linea)
		os.Exit(1)
	}

}

func validarErroresNumericos(contenido []rune, i int, mejorLongitud int, mejorLexema *Token, linea int) {
	if mejorLexema.NombreToken != "token_entero" && mejorLexema.NombreToken != "token_flotante" {
		return
	}

	n := len(contenido)
	posSig := i + mejorLongitud

	if posSig < n && esLetraOSimboloId(contenido[posSig]) {
		fmt.Printf("Error lexico: identificador no puede comenzar con digito en linea %d\n", linea)
		os.Exit(1)
	}

	if posSig < n && contenido[posSig] == '.' {
		if posSig+1 >= n || contenido[posSig+1] != '.' {
			fmt.Printf("Error lexico: literal flotante mal formado en linea %d\n", linea)
			os.Exit(1)
		}
	}
}
