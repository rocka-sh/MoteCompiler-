package literales

import (
	adf "MoteCompiler/analizadorlexico"
)

var Literales = adf.Tokenario{
	Tipo: "literales",
	Tokens: []*adf.Token{
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
