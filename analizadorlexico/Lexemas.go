package analizadorlexico

type Lexema struct {
	QInicial Estado
	Estados  []Estado
	Token    string
}

func CrearLexema(estadoInicial Estado, listaEstados []Estado, nombreToken string) *Lexema {
	var l Lexema
	l.QInicial = estadoInicial
	l.Estados = listaEstados
	l.Token = nombreToken
	return &l
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
