package main

import (
	adf "MoteCompiler/analizadorlexico"
	"MoteCompiler/analizadorsintactico/parser"
	"fmt"
	"os"
)

func Sintactico(tokens []adf.Token) {

	p := &parser.Parser{
		Tokens:     tokens,
		Cursor:     0,
		NumErrores: 0,
	}

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(parser.SyntaxError); !ok {
				panic(r)
			}
		}
		if p.NumErrores > 0 {
			fmt.Printf("\n se encontraron %d errores sintácticos.\n", p.NumErrores)
			os.Exit(1)
		}
	}()

	p.ParseProgram()
	p.Match("$")

	if p.NumErrores > 0 {
		fmt.Printf("\nse encontraron %d errores sintácticos.\n", p.NumErrores)
		os.Exit(1)
	}

	fmt.Println("el codigo es correcto")
}
