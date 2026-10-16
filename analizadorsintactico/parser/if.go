package parser

func (p *Parser) ParseIfStmt() {
	p.Match("if")
	p.ParseExpr()
	p.Match("then")
	p.ParseBlock()
	p.ParseElifList()
	p.ParseElseOpt()
	p.Match("end")
}

func (p *Parser) ParseElifList() {
	switch p.tokenActual().NombreToken {
	case "elif":
		p.ParseElifClause()
		p.ParseElifList()
	case "else", "end":
		// e (epsilon)
		return
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseElifClause() {
	p.Match("elif")
	p.ParseExpr()
	p.Match("then")
	p.ParseBlock()
}

func (p *Parser) ParseElseOpt() {
	switch p.tokenActual().NombreToken {
	case "else":
		p.Match("else")
		p.ParseBlock()
	case "end":
		// e (epsilon)
		return
	default:
		p.errorInesperado()
	}
}
