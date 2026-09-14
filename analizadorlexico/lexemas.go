package analizadorlexico

type Lexema struct {
	QInicial Estado
	Token    string
}

func (l *Lexema) D(r []rune) int {
	q := &l.QInicial

	var longitudMax int = 0
	var acumulado int = 0

	for _, v := range r {
		siguienteEstado := q.d(v)

		//recibio un caracter no aceptado en el adf
		if siguienteEstado == nil {
			break
		}
		acumulado++
		q = siguienteEstado

		if q.IsF {
			longitudMax = acumulado
		}
	}

	return longitudMax
}
