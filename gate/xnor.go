package gate

type xnor struct {
	gate
}

func (g xnor) Output() (bool, error) {
	switch {
	case g.a == true && g.b == true:
		return true, nil
	case g.a == false && g.b == false:
		return true, nil
	default:
		return false, nil
	}
}
