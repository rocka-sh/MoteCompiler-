package analizadorlexico

type Token struct {
	QInicial    Estado
	NombreToken string
	Lexema      string
	Linea       int
}

func (l *Token) D(r []rune) int {
	q := &l.QInicial

	var longitudMax int = 0
	var acumulado int = 0

	for _, v := range r {
		siguienteEstado := q.d(v)

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
