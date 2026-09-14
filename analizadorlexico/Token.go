package analizadorlexico

type Token struct {
	Tipo    string
	Lexemas []*Lexema
}

func (t Token) EvaluarPrefijo(r []rune) (*Lexema, int) {
	var mejorLongitud int = 0
	var mejorLexema *Lexema = nil

	for _, lexema := range t.Lexemas {
		var longitud int
		longitud = lexema.D(r)

		if longitud > mejorLongitud {
			mejorLongitud = longitud
			mejorLexema = lexema
		}
	}

	return mejorLexema, mejorLongitud
}
