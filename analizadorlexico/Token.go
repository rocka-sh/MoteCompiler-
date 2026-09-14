package analizadorlexico

import (
	"fmt"
)

type Token struct {
	Tipo    string
	Lexemas []Lexema
}

func (t Token) Tokenizador(runes []rune) {
	for _, lexema := range t.Lexemas {
		if lexema.D(runes) != nil {
			fmt.Println(t.Tipo, " ", string(runes))
			break
		}
	}
}
