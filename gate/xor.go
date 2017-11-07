package gate

type xor struct {
	gate
}

func (g xor) Output() (bool, error) {
	switch {
	case g.a == true && g.b == true:
		return false, nil
	case g.a == false && g.b == false:
		return false, nil
	default:
		return true, nil
	}
}
