package parser

// Reglas en gramática:
// Var = VarKw Id ":" Type VarAssig
// VarKw = "let" | "var"
// VarAssig = ":=" Expr | e

func (p *Parser) ParseVar() {
	p.ParseVarKw()
	p.Match("Id")
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
