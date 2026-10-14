package parser

// Reglas en gramática:
// Expr = OrExpr
// OrExpr = AndExpr OrExprT
// OrExprT = "or" AndExpr OrExprT | e
// AndExpr = NotExpr AndExprT
// AndExprT = "and" NotExpr AndExprT | e
// NotExpr = "not" NotExpr | Comparison
// Comparison = Additive ComparisonOpt
// ComparisonOpt = CompOp Additive | e
// CompOp = "=" | "<>" | "<" | "<=" | ">" | ">="
// Additive = Multiplicative AdditiveT
// AdditiveT = AddOp Multiplicative AdditiveT | e
// AddOp = "+" | "-"
// Multiplicative = Unary MultiplicativeT
// MultiplicativeT = MultOp Unary MultiplicativeT | e
// MultOp = "*" | "/" | "%"
// Unary = "-" Unary | Primary
// Primary = IntLiteral | FloatLiteral | StringLiteral | "true" | "false" | PrimaryIdent | ArrayLiteral | "(" Expr ")"
// PrimaryIdent = Id PrimaryIdentOpt
// PrimaryIdentOpt = "(" OptArgList ")" | ValueList
// ArrayLiteral = "[" OptArgList "]"
// OptArgList = ArgList | e
// ArgList = Expr ArgListT
// ArgListT = "," Expr ArgListT | e

func (p *Parser) ParseExpr() {
	p.ParseOrExpr()
}

func (p *Parser) ParseOrExpr() {
	p.ParseAndExpr()
	p.ParseOrExprT()
}

func (p *Parser) ParseOrExprT() {
	switch p.tokenActual().NombreToken {
	case "or":
		p.Match("or")
		p.ParseAndExpr()
		p.ParseOrExprT()
	default:
		// e (epsilon)
		return
	}
}

func (p *Parser) ParseAndExpr() {
	p.ParseNotExpr()
	p.ParseAndExprT()
}

func (p *Parser) ParseAndExprT() {
	switch p.tokenActual().NombreToken {
	case "and":
		p.Match("and")
		p.ParseNotExpr()
		p.ParseAndExprT()
	default:
		// e (epsilon)
		return
	}
}

func (p *Parser) ParseNotExpr() {
	switch p.tokenActual().NombreToken {
	case "not":
		p.Match("not")
		p.ParseNotExpr()
	default:
		p.ParseComparison()
	}
}

func (p *Parser) ParseComparison() {
	p.ParseAdditive()
	p.ParseComparisonOpt()
}

func (p *Parser) ParseComparisonOpt() {
	switch p.tokenActual().NombreToken {
	case "=", "<>", "<", "<=", ">", ">=":
		p.ParseCompOp()
		p.ParseAdditive()
	default:
		// e (epsilon)
		return
	}
}

func (p *Parser) ParseCompOp() {
	switch p.tokenActual().NombreToken {
	case "=", "<>", "<", "<=", ">", ">=":
		p.Match(p.tokenActual().NombreToken)
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseAdditive() {
	p.ParseMultiplicative()
	p.ParseAdditiveT()
}

func (p *Parser) ParseAdditiveT() {
	switch p.tokenActual().NombreToken {
	case "+", "-":
		p.ParseAddOp()
		p.ParseMultiplicative()
		p.ParseAdditiveT()
	default:
		// e (epsilon)
		return
	}
}

func (p *Parser) ParseAddOp() {
	switch p.tokenActual().NombreToken {
	case "+", "-":
		p.Match(p.tokenActual().NombreToken)
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseMultiplicative() {
	p.ParseUnary()
	p.ParseMultiplicativeT()
}

func (p *Parser) ParseMultiplicativeT() {
	switch p.tokenActual().NombreToken {
	case "*", "/", "%":
		p.ParseMultOp()
		p.ParseUnary()
		p.ParseMultiplicativeT()
	default:
		// e (epsilon)
		return
	}
}

func (p *Parser) ParseMultOp() {
	switch p.tokenActual().NombreToken {
	case "*", "/", "%":
		p.Match(p.tokenActual().NombreToken)
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseUnary() {
	switch p.tokenActual().NombreToken {
	case "-":
		p.Match("-")
		p.ParseUnary()
	default:
		p.ParsePrimary()
	}
}

func (p *Parser) ParsePrimary() {
	switch p.tokenActual().NombreToken {
	case "IntLiteral":
		p.Match("IntLiteral")
	case "FloatLiteral":
		p.Match("FloatLiteral")
	case "StringLiteral":
		p.Match("StringLiteral")
	case "true":
		p.Match("true")
	case "false":
		p.Match("false")
	case "Id":
		p.ParsePrimaryIdent()
	case "[":
		p.ParseArrayLiteral()
	case "(":
		p.Match("(")
		p.ParseExpr()
		p.Match(")")
	case "+", "*", "/", "%", "=", "<>", "<", "<=", ">", ">=", "and", "or":
		p.errorSintactico("operador duplicado o falta de operando en la expresión", p.tokenActual().Linea)
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParsePrimaryIdent() {
	p.Match("Id")
	p.ParsePrimaryIdentOpt()
}

func (p *Parser) ParsePrimaryIdentOpt() {
	switch p.tokenActual().NombreToken {
	case "(":
		p.Match("(")
		p.ParseOptArgList()
		p.Match(")")
	default:
		p.ParseValueList()
	}
}

func (p *Parser) ParseArrayLiteral() {
	p.Match("[")
	p.ParseOptArgList()
	p.Match("]")
}

func (p *Parser) ParseOptArgList() {
	switch p.tokenActual().NombreToken {
	case "(", "-", "[", "false", "not", "true", "FloatLiteral", "Id", "IntLiteral", "StringLiteral":
		p.ParseArgList()
	case ")", "]":
		// e (epsilon)
		return
	default:
		p.errorInesperado()
	}
}

func (p *Parser) ParseArgList() {
	p.ParseExpr()
	p.ParseArgListT()
}

func (p *Parser) ParseArgListT() {
	switch p.tokenActual().NombreToken {
	case ",":
		p.Match(",")
		if p.tokenActual().NombreToken == ")" || p.tokenActual().NombreToken == "]" {
			p.errorSintactico("coma final sobrante en lista de argumentos", p.tokenActual().Linea)
		}
		p.ParseExpr()
		p.ParseArgListT()
	case ")", "]":
		// e (epsilon)
		return
	default:
		p.errorInesperado()
	}
}
