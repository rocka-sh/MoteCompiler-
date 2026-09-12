package analizadorlexico

type Lexema struct {
	Q1      Estado
	Estados []Estado
	Token   string
}

func (l *Lexema) D(r []rune) *Lexema {
	qActual := &l.Q1

	for _, v := range r {
		siguienteEstado := qActual.d(v)

		//recibio un caracter no aceptado en el adf
		if siguienteEstado == nil {
			return nil
		}

		//no fue un estado aceptado y faltan caracteres por evaluar
		qActual = siguienteEstado
	}

	//se acabaron los caracteres pero termino en un estado F aceptado
	if qActual.IsF {
		return l
	}

	//se acabaron los caracteres y no termino en un Estado Aceptado
	return nil
}

type Estado struct {
	CharsAceptados map[*Estado][]rune
	IsF            bool
}

func (e *Estado) d(w rune) *Estado {
	for q, v := range e.CharsAceptados {
		for _, i := range v {
			if w == i {
				return q
			}
		}
	}
	return nil
}
