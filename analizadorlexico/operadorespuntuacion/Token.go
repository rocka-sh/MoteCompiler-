package operadorespuntuacion

import (
	adf "MoteCompiler/analizadorlexico"
)

var OperadoresPuntuacion = adf.Token{
	Tipo: "Operador/Puntuacion",
	Lexemas: []*adf.Lexema{
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
