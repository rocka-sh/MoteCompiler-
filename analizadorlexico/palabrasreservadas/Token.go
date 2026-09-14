package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var PalabrasReservadas = adf.Token{
	Tipo: "Palabra Reservada",
	Lexemas: []adf.Lexema{
		LexemaAnd,
		LexemaBool,
		LexemaInt,
		LexemaFloat,
		LexemaString,
		LexemaOr,
		LexemaNot,
		LexemaIf,
		LexemaThen,
		LexemaElse,
		LexemaElif,
		LexemaEnd,
		LexemaWhile,
		LexemaFor,
		LexemaDo,
		LexemaReturn,
		LexemaFn,
		LexemaIn,
		LexemaPrint,
		LexemaRecord,
		LexemaLet,
		LexemaVar,
	},
}
