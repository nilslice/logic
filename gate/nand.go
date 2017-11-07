package gate

type nand struct {
	gate
}

func (g nand) Output() (bool, error) {
	switch {
	case g.a == true && g.b == true:
		return false, nil
	default:
		return true, nil
	}
}
