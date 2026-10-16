package parser

func (p *Parser) ParseVar() {
	p.ParseVarKw()
	p.Match("Id")
	if p.tokenActual().NombreToken == ":=" || p.tokenActual().NombreToken == "=" {
		p.errorSintactico("falta el tipo de dato con ':' antes de la asignación", p.tokenActual().Linea)
	}
	p.Match(":")
	p.ParseType()
	p.ParseVarAssig()
}

func (p *Parser) ParseVarKw() {
	switch p.tokenActual().NombreToken {
	case "let":
		p.Match("let")
	case "var":
		p.Match("var")
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseVarAssig() {
	switch p.tokenActual().NombreToken {
	case ":=":
		p.Match(":=")
		p.ParseExpr()
	default:
		// e (epsilon)
		return
	}
}
