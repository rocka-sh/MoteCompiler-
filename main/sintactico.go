package main

import (
	adf "MoteCompiler/analizadorlexico"
	"MoteCompiler/analizadorsintactico/parser"
	"fmt"
)

func Sintactico(tokens []adf.Token) {

	p := &parser.Parser{
		Tokens: tokens,
		Cursor: 0,
	}
	p.ParseProgram()
	p.Match("$")
	fmt.Println("el codigo es correcto")
}
