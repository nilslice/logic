package gate

type not struct {
	gate
}

func (g not) Output() (bool, error) {
	return !g.a, nil
}
