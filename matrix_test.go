package matrix

import (
	"testing"
)

//Test the identity matrix creation
func TestIdentity(t *testing.T) {
	cases := []struct {
		in   int
		want Matrix
	}{
		{in: 2, want: Matrix{[][]float64{{1.0, 0.0}, {0.0, 1.0}}, 2, 2}},
		{in: 3, want: Matrix{[][]float64{{1.0, 0.0, 0.0}, {0.0, 1.0, 0.0}, {0.0, 0.0, 1.0}}, 2, 2}},
	}
	for _, c := range cases {
		got := I(c.in)
		if got.n != c.want.n || got.m != c.want.m {
			t.Errorf("Dim. err. want: %v x %v got: $v x $v", c.want.m, c.want.n, got.m, got.n)
		}
		if !equal(c.want, got) {
			t.Errorf("Matrix elements do not match %v || %v ", c.want.matrix, got.matrix)
		}
	}
}
