package gate

type nor struct {
	gate
}

func (g nor) Output() (bool, error) {
	switch {
	case g.a == false && g.b == false:
		return true, nil
	default:
		return false, nil
	}
}
