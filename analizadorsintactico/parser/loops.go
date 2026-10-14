package parser

// Reglas en gramática:
// WhileStmt = "while" Expr "do" Block "end"
// ForStmt = "for" Id "in" Expr ".." Expr "do" Block "end"

func (p *Parser) ParseWhileStmt() {
	p.Match("while")
	p.ParseExpr()
	p.Match("do")
	p.ParseBlock()
	p.Match("end")
}

func (p *Parser) ParseForStmt() {
	p.Match("for")
	p.Match("Id")
	p.Match("in")
	p.ParseExpr()
	p.Match("..")
	p.ParseExpr()
	p.Match("do")
	p.ParseBlock()
	p.Match("end")
}
