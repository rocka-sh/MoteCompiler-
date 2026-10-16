package parser

func (p *Parser) ParseProgram() {
	p.ParseDeclarationList()
}

func (p *Parser) ParseDeclarationList() {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(SyntaxError); ok {
				p.ParseDeclarationList()
				return
			}
			panic(r)
		}
	}()

	switch p.tokenActual().NombreToken {
	case "fn", "let", "var", "record":
		p.ParseDeclaration()
		p.ParseDeclarationList()
	case "$":
		// e (epsilon)
		return
	case "if", "while", "for", "return", "print":
		p.errorSintactico("las sentencias ejecutables deben estar dentro de una funcion", p.tokenActual().Linea)
		p.ParseDeclarationList()
	default:
		p.errorInesperado()
		p.ParseDeclarationList()
	}
}

func (p *Parser) ParseDeclaration() {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(SyntaxError); ok {
				return
			}
			panic(r)
		}
	}()

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
