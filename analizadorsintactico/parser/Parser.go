package parser

import (
	adf "MoteCompiler/analizadorlexico"
	"fmt"
)

type SyntaxError struct{}

type Parser struct {
	Tokens     []adf.Token
	Cursor     int
	NumErrores int
}

func (p *Parser) tokenActual() adf.Token {
	if p.Cursor < len(p.Tokens) {
		return p.Tokens[p.Cursor]
	}
	return adf.Token{NombreToken: "$", Lexema: "", Linea: -1}
}

func (p *Parser) errorSintactico(mssg string, linea int) {
	fmt.Printf("Error sintáctico en línea %d: %s\n", linea, mssg)
	p.NumErrores++
	p.sincronizar()
	panic(SyntaxError{})
}

func (p *Parser) errorInesperado() {
	tok := p.tokenActual()
	p.errorSintactico("token inesperado '"+tok.NombreToken+"' (lexema: '"+tok.Lexema+"')", tok.Linea)
}

func (p *Parser) sincronizar() {
	p.Cursor++
	for p.tokenActual().NombreToken != "$" {
		switch p.tokenActual().NombreToken {
		case "end", "if", "while", "for", "let", "var", "return", "print", "Id", "fn", "record":
			return
		}
		p.Cursor++
	}
}

func (p *Parser) Match(esperado string) {
	if p.tokenActual().NombreToken == esperado {
		p.Cursor++
	} else {
		p.errorSintactico("se esperaba '"+esperado+"' pero se encontró '"+p.tokenActual().NombreToken+"'", p.tokenActual().Linea)
	}
}
