package palabrasreservadas

import (
	adf "MoteCompiler/analizadorlexico"
)

var PalabrasReservadas = adf.Tokenario{
	Tipo: "palabra_reservada",
	Tokens: []*adf.Token{
		&LexemaAnd,
		&LexemaBool,
		&LexemaInt,
		&LexemaFloat,
		&LexemaString,
		&LexemaOr,
		&LexemaNot,
		&LexemaIf,
		&LexemaThen,
		&LexemaElse,
		&LexemaElif,
		&LexemaEnd,
		&LexemaWhile,
		&LexemaFor,
		&LexemaDo,
		&LexemaReturn,
		&LexemaFn,
		&LexemaIn,
		&LexemaPrint,
		&LexemaRecord,
		&LexemaLet,
		&LexemaVar,
	},
}
