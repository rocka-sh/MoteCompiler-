package parser

// Reglas en gramática:
// Type = "int" | "float" | "string" | "bool" | ArrayType | Id
// ArrayType = "[" IntLiteral "]" Type

func (p *Parser) ParseType() {
	switch p.tokenActual().NombreToken {
	case "int":
		p.Match("int")
	case "float":
		p.Match("float")
	case "string":
		p.Match("string")
	case "bool":
		p.Match("bool")
	case "[":
		p.ParseArrayType()
	case "Id":
		p.Match("Id")
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseArrayType() {
	p.Match("[")
	p.Match("IntLiteral")
	p.Match("]")
	p.ParseType()
}
