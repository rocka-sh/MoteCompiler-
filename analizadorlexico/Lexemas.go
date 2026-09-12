package analizadorlexico

type Lexema struct {
	QInicial Estado
	Token    string
}

func (l *Lexema) D(r []rune) *Lexema {
	q := &l.QInicial
	for _, v := range r {
		siguienteEstado := q.d(v)

		//recibio un caracter no aceptado en el adf
		if siguienteEstado == nil {
			return nil
		}

		q = siguienteEstado
	}

	if q.IsF {
		return l
	}

	return nil
}
