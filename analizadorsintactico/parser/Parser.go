package parser

import (
	adf "MoteCompiler/analizadorlexico"
	"fmt"
	"os"
)

type Parser struct {
	Tokens []adf.Token
	Cursor int
}

func (p *Parser) tokenActual() adf.Token {
	return p.Tokens[p.Cursor]
}

func (p *Parser) errorSintactico(mssg string, linea int) {
	fmt.Printf("Error sintáctico en línea %d: %s\n", linea, mssg)
	os.Exit(1)
}

func (p *Parser) errorInesperado() {
	tok := p.tokenActual()
	fmt.Printf("Error sintáctico en línea %d: token inesperado '%s' (lexema: '%s')\n",
		tok.Linea, tok.NombreToken, tok.Lexema)
	os.Exit(1)
}

func (p *Parser) Match(esperado string) {
	if p.tokenActual().NombreToken == esperado {
		p.Cursor++
	} else {
		p.errorSintactico("se esperaba '"+esperado+"' pero se encontró '"+p.tokenActual().NombreToken+"'", p.tokenActual().Linea)
	}
}
