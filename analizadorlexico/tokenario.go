package analizadorlexico

type Tokenario struct {
	Tipo   string
	Tokens []*Token
}

func (t Tokenario) EvaluarPrefijo(r []rune) (*Token, int) {
	var mejorLongitud int = 0
	var mejorLexema *Token = nil

	for _, lexema := range t.Tokens {
		var longitud int
		longitud = lexema.D(r)

		if longitud > mejorLongitud {
			mejorLongitud = longitud
			mejorLexema = lexema
		}
	}

	return mejorLexema, mejorLongitud
}
