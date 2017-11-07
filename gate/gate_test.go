package gate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAndGate(t *testing.T) {
	cases := []struct {
		name      string
		a, b, exp bool
	}{
		{"0, 0 = 0", false, false, false},
		{"0, 1 = 0", false, true, false},
		{"1, 0 = 0", true, false, false},
		{"1, 1 = 1", true, true, true},
	}

	for _, tc := range cases {
		g := and{}
		g.Input(tc.a, tc.b)
		y, err := g.Output()
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, y)
	}
}

func TestNotGate(t *testing.T) {
	cases := []struct {
		name   string
		a, exp bool
	}{
		{"0 = 1", true, false},
		{"1 = 0", false, true},
	}

	for _, tc := range cases {
		g := not{}
		g.Input(tc.a)
		y, err := g.Output()
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, y)
	}
}

func TestOrGate(t *testing.T) {
	cases := []struct {
		name      string
		a, b, exp bool
	}{
		{"0, 0 = 0", false, false, false},
		{"0, 1 = 1", false, true, true},
		{"1, 0 = 1", true, false, true},
		{"1, 1 = 1", true, true, true},
	}

	for _, tc := range cases {
		g := or{}
		g.Input(tc.a, tc.b)
		y, err := g.Output()
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, y)
	}
}

func TestNorGate(t *testing.T) {
	cases := []struct {
		name      string
		a, b, exp bool
	}{
		{"1, 0 = 0", true, false, false},
		{"0, 1 = 0", false, true, false},
		{"1, 1 = 0", true, true, false},
		{"0, 0 = 1", false, false, true},
	}

	for _, tc := range cases {
		g := nor{}
		g.Input(tc.a, tc.b)
		y, err := g.Output()
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, y)
	}
}

func TestNandGate(t *testing.T) {
	cases := []struct {
		name      string
		a, b, exp bool
	}{
		{"0, 1 = 1", false, true, true},
		{"1, 0 = 1", true, false, true},
		{"0, 0 = 1", false, false, true},
		{"1, 1 = 0", true, true, false},
	}

	for _, tc := range cases {
		g := nand{}
		g.Input(tc.a, tc.b)
		y, err := g.Output()
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, y)
	}
}

func TestXorGate(t *testing.T) {
	cases := []struct {
		name      string
		a, b, exp bool
	}{
		{"0, 0 = 0", false, false, false},
		{"0, 1 = 1", false, true, true},
		{"1, 0 = 1", true, false, true},
		{"1, 1 = 0", false, false, false},
	}

	for _, tc := range cases {
		g := xor{}
		g.Input(tc.a, tc.b)
		y, err := g.Output()
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, y)
	}
}

func TestXnorGate(t *testing.T) {
	cases := []struct {
		name      string
		a, b, exp bool
	}{
		{"0, 0 = 1", false, false, true},
		{"0, 1 = 0", false, true, false},
		{"1, 0 = 0", true, false, false},
		{"1, 1 = 1", true, true, true},
	}

	for _, tc := range cases {
		g := xnor{}
		g.Input(tc.a, tc.b)
		y, err := g.Output()
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, y)
	}
}
