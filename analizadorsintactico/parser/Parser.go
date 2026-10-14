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

func (p *Parser) Match(esperado string) {
	if p.tokenActual().NombreToken == esperado {
		p.Cursor++
	} else {
		fmt.Printf("Error sintáctico en línea %d: se esperaba '%s' pero se encontró '%s'\n",
			p.tokenActual().Linea, esperado, p.tokenActual().NombreToken)
		os.Exit(1)
	}
}

func (p *Parser) errorInesperado() {
	tok := p.tokenActual()
	fmt.Printf("Error sintáctico en línea %d: token inesperado '%s' (lexema: '%s')\n",
		tok.Linea, tok.NombreToken, tok.Lexema)
	os.Exit(1)
}
