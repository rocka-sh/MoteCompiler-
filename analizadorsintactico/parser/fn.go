package parser

func (p *Parser) ParseFn() {
	p.Match("fn")
	p.Match("Id")
	p.Match("(")
	p.ParseParamList()
	p.Match(")")
	p.ParseTypeFun()
	p.ParseBlock()
	p.Match("end")
}

func (p *Parser) ParseTypeFun() {
	switch p.tokenActual().NombreToken {
	case ":":
		p.Match(":")
		p.ParseType()
	default:
		// e (epsilon)
		return
	}
}

func (p *Parser) ParseParamList() {
	switch p.tokenActual().NombreToken {
	case "Id":
		p.ParseParam()
		p.ParseParamT()
	case ")":
		// e (epsilon)
		return
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseParamT() {
	switch p.tokenActual().NombreToken {
	case ",":
		p.Match(",")
		if p.tokenActual().NombreToken == ")" {
			p.errorSintactico("coma final sobrante en lista de parametros", p.tokenActual().Linea)
		}
		p.ParseParam()
		p.ParseParamT()
	case ")":
		// e (epsilon)
		return
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseParam() {
	p.Match("Id")
	p.Match(":")
	p.ParseType()
}
