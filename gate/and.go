package gate

type and struct {
	gate
}

func (g and) Output() (bool, error) {
	switch {
	case g.a == true && g.b == true:
		return true, nil
	default:
		return false, nil
	}
}
