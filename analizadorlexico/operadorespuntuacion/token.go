package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

var OperadoresPuntuacion = adf.Tokenario{
	Tipo: "operador/puntuacion",
	Tokens: []*adf.Token{
		&LexemaAsignacion,
		&LexemaIgual,
		&LexemaComparacion,
		&LexemaMenor,
		&LexemaMenorIgual,
		&LexemaMayor,
		&LexemaMayorIgual,
		&LexemaSuma,
		&LexemaResta,
		&LexemaMultiplicacion,
		&LexemaDivision,
		&LexemaModulo,
		&LexemaParentesisAbre,
		&LexemaParentesisCierra,
		&LexemaCorcheteAbre,
		&LexemaCorcheteCierra,
		&LexemaLlaveAbre,
		&LexemaLlaveCierra,
		&LexemaComa,
		&LexemaDosPuntos,
		&LexemaRango,
	},
}
