package analizadorlexico

// TODO: funciones aux para los estados sin tranciciones
type Estado struct {
	Transiciones map[rune]*Estado
	IsF          bool
}

func (e *Estado) d(w rune) *Estado {
	var siguiente *Estado = e.Transiciones[w]
	return siguiente
}
