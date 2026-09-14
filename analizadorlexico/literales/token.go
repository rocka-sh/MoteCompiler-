package literales

import (
	adf "MoteCompiler/analizadorlexico"
)

var Literales = adf.Token{
	Tipo: "literales",
	Lexemas: []*adf.Lexema{
		&LexemaTrue,
		&LexemaFalse,
		&LexemaEntero,
		&LexemaFlotante,
		&LexemaCadena,
	},
}

func InitADF() {
	initNumeros()
	initCadenas()
}
