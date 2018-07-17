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
		{in: 3, want: Matrix{[][]float64{{1.0, 0.0, 0.0},
			{0.0, 1.0, 0.0}, {0.0, 0.0, 1.0}}, 2, 2}},
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

func TestQR(t *testing.T) {
	cases := []struct {
		in    Matrix
		wantR Matrix
		wantQ Matrix
	}{
		{in: Matrix{[][]float64{{1.0, 1.0, 2.0}, {1.0, 0.0, -2.0}, {-1.0, 2.0, 3.0}}, 3, 3},
			wantR: Matrix{[][]float64{{1.7321, -0.5774, -1.7321},
				{0.0, 2.1603, 3.2404}, {0.0, 0.0, 1.8708}}, 3, 3},
			wantQ: Matrix{[][]float64{{0.5574, 0.6172, 0.5345},
				{0.5574, 0.1543, -0.8018}, {-0.5574, 0.7715, -0.2673}}, 3, 3},
		},
	}
	for _, c := range cases {
		gotQ, gotR := QR(c.in)
		if gotR.n != c.wantR.n || gotR.m != c.wantR.m || gotQ.n != c.wantQ.n || gotQ.m != c.wantQ.m {
			t.Errorf("Dim. err. want: %v x %v got: $v x $v & $v x $v", c.wantR.m, c.wantR.n,
				c.wantQ.m, c.wantQ.n, gotQ.m, gotQ.n)
		}
		if !equal(c.wantR, gotR) {
			t.Errorf("Matrix elements of R do not match %v || %v ", c.wantR.matrix, gotR.matrix)
		}
		if !equal(c.wantQ, gotQ) {
			t.Errorf("Matrix elements of Q do not match %v || %v ", c.wantQ.matrix, gotQ.matrix)
		}
	}
}
