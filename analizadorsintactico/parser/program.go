package parser

// Reglas en gramática:
// Program = DeclarationList
// DeclarationList = Declaration DeclarationList | e
// Declaration = Fn | Var | Record

func (p *Parser) ParseProgram() {
	p.ParseDeclarationList()
}

func (p *Parser) ParseDeclarationList() {
	switch p.tokenActual().NombreToken {
	case "fn", "let", "var", "record":
		p.ParseDeclaration()
		p.ParseDeclarationList()
	case "$":
		// e (epsilon)
		return
	case "if", "while", "for", "return", "print":
		p.errorSintactico("las sentencias ejecutables deben estar dentro de una función", p.tokenActual().Linea)
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseDeclaration() {
	switch p.tokenActual().NombreToken {
	case "fn":
		p.ParseFn()
	case "let", "var":
		p.ParseVar()
	case "record":
		p.ParseRecord()
	default:
		p.errorInesperado()
	}
}
