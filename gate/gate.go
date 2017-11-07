package gate

// Gate describes a logic gate
type Gate interface {
	Input(inputs ...bool)
	Output() (y bool, err error)
}

type gate struct {
	a, b, y bool
}

func (g *gate) Input(inputs ...bool) {
	g.a = inputs[0]
	if len(inputs) > 1 {
		g.b = inputs[1]
	}
}
