package parser

func (p *Parser) ParseBlock() {
	p.ParseStatementList()
}

func (p *Parser) ParseStatementList() {
	switch p.tokenActual().NombreToken {
	case "let", "var", "if", "while", "for", "return", "print", "Id":
		p.ParseStatement()
		p.ParseStatementList()
	case "elif", "else", "end":
		// e (epsilon)
		return
	case "$":
		p.errorSintactico("fin de archivo inesperado, falta un 'end'", p.tokenActual().Linea)
		return
	default:
		p.errorInesperado()
		p.ParseStatementList()
	}
}

func (p *Parser) ParseStatement() {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(SyntaxError); ok {
				return
			}
			panic(r)
		}
	}()

	switch p.tokenActual().NombreToken {
	case "let", "var":
		p.ParseVar()
	case "if":
		p.ParseIfStmt()
	case "while":
		p.ParseWhileStmt()
	case "for":
		p.ParseForStmt()
	case "return":
		p.ParseReturnStmt()
	case "print":
		p.ParsePrintStmt()
	case "Id":
		p.ParseStmtStartId()
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseStmtStartId() {
	p.Match("Id")
	p.ParseStmtIdTail()
}

func (p *Parser) ParseStmtIdTail() {
	switch p.tokenActual().NombreToken {
	case "(":
		p.Match("(")
		p.ParseOptArgList()
		p.Match(")")
	case ".", "[", ":=":
		p.ParseValueList()
		p.Match(":=")
		p.ParseExpr()
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseValueList() {
	switch p.tokenActual().NombreToken {
	case ".", "[":
		p.ParseValue()
		p.ParseValueList()
	default:
		// e (epsilon)
		return
	}
}

func (p *Parser) ParseValue() {
	switch p.tokenActual().NombreToken {
	case "[":
		p.Match("[")
		p.ParseExpr()
		p.Match("]")
	case ".":
		p.Match(".")
		p.Match("Id")
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseReturnStmt() {
	p.Match("return")
	p.ParseExpr()
}

func (p *Parser) ParsePrintStmt() {
	p.Match("print")
	p.Match("(")
	p.ParseExpr()
	p.Match(")")
}
