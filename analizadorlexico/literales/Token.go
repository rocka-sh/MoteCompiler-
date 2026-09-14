package literales

import (
	adf "MoteCompiler/analizadorlexico"
)

var Literales = adf.Token{
	Tipo: "Literales",
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
