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
	var i int = 0
	var n int = len(contenido)

	for i < n {
		if unicode.IsSpace(contenido[i]) {
			i++
			continue
		}

		var subslice []rune = contenido[i:]
		var lexemaEncontrado *adf.Lexema = nil
		var longitud int = 0

		//primero palabras reservadas
		lexemaEncontrado, longitud = tokenReservadas.EvaluarPrefijo(subslice)

		// Literales (si no fue palabra reservada)
		if longitud == 0 {
			lexemaEncontrado, longitud = tokenLiterales.EvaluarPrefijo(subslice)
		}

		// Identificadores (si no fue literal ni reservada)
		if longitud == 0 {
			lexemaEncontrado, longitud = tokenIdentificador.EvaluarPrefijo(subslice)
		}

		if longitud > 0 {
			var textoToken string = string(contenido[i : i+longitud])
			fmt.Println("Token:", lexemaEncontrado.Token, "Lexema:", textoToken)
			i += longitud
		} else {
			fmt.Println("Caracteres no reconocidos: ", string(contenido[i]))
			i++
		}
	}
}
