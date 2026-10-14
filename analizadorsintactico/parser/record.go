package parser

// Reglas en gramática:
// Record = "record" Id FieldList "end"
// FieldList = FieldDecl FieldList | e
// FieldDecl = Id ":" Type

func (p *Parser) ParseRecord() {
	p.Match("record")
	p.Match("Id")
	p.ParseFieldList()
	p.Match("end")
}

func (p *Parser) ParseFieldList() {
	switch p.tokenActual().NombreToken {
	case "Id":
		p.ParseFieldDecl()
		p.ParseFieldList()
	case "end":
		// e (epsilon)
		return
	case "$":
		p.errorSintactico("fin de archivo inesperado, falta un 'end' para cerrar el record", p.tokenActual().Linea)
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseFieldDecl() {
	p.Match("Id")
	p.Match(":")
	p.ParseType()
}
