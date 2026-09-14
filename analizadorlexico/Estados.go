package analizadorlexico

type Estado struct {
	Transiciones map[rune]*Estado
	IsF          bool
}

func (e *Estado) d(w rune) *Estado {
	if e.Transiciones == nil {
		return nil
	}
	var siguiente *Estado = e.Transiciones[w]
	return siguiente
}
