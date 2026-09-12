package analizadorlexico

type Lexema struct {
	Q1      Estado
	Estados []Estado
	Token   string
}

func CrearLexema(estadoInicial Estado, listaEstados []Estado, nombreToken string) *Lexema {
	var l Lexema
	l.Q1 = estadoInicial
	l.Estados = listaEstados
	l.Token = nombreToken
	return &l
}

// TODO: funciones aux para los estados sin tranciciones
func (l *Lexema) D(r []rune) *Lexema {
	qActual := &l.Q1
	for _, v := range r {
		siguienteEstado := qActual.d(v)

		//recibio un caracter no aceptado en el adf
		if siguienteEstado == nil {
			return nil
		}

		qActual = siguienteEstado
	}

	//termino en un estado F aceptado
	if qActual.IsF {
		return l
	}
	//no termino en un Estado Aceptado
	return nil
}
