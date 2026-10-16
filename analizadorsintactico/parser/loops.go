package parser

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
