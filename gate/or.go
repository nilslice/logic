package gate

type or struct {
	gate
}

func (g or) Output() (bool, error) {
	switch {
	case g.a == false && g.b == false:
		return false, nil
	default:
		return true, nil
	}
}
