package matrix

import (
	"fmt"
	"math"
	"math/rand"
)

/*
A simple matrix operations package for statistical procedures
Implements a matrix type
*/

//Defning a simple Tensor interface
type Tensor interface {
	//returns the matrix dimensions
	//width and height
	dim() (width, height int)
	// returns the complex conjugate
	// of the matrix
	cc() Matrix
	//returns transpose of the matrix
	t() Matrix
	//returns a nxn identity matrix
	I(n int) Matrix
	//Prints an rperesentation of the matrix
	Print() string
	//returns an m x m matrix with random floats
	genRandMatrix(m, n int) Matrix
}

//Matrix of size m x n
type Matrix struct {
	matrix [][]float64
	m      int //dimensions n
	n      int
}

//generates a n X m matrix fillled
//with random floats
func genRandMatrix(m, n int) Matrix {
	//if dims less than zero panic
	if n < 0 || m < 0 {
		panic(fmt.Sprintf("m or n is less than zero n=%v, m=%v", m, n))
	}
	//create 2d matrix
	retArr := create2DSlice(m, n)
	//fill with random values
	for i, elem := range retArr {
		for j, _ := range elem {
			retArr[i][j] = rand.Float64()
		}
	}
	return Matrix{retArr, m, n}
}

//generates a n X m matrix fillled
//with random floats
func I(n int) Matrix {
	//if dims less than zero panic
	if n < 0 {
		panic(fmt.Sprintf("n is less than zero n=%v", n))
	}
	//create 2d matrix
	retArr := create2DSlice(n, n)
	//fill with 1s along diag
	for i := 0; i < n; i++ {
		retArr[i][i] = 1.0
	}
	return Matrix{retArr, n, n}
}

//Helper function creates an empty 2d slice
func create2DSlice(m, n int) [][]float64 {
	retArr := make([][]float64, m)
	for row := range retArr {
		retArr[row] = make([]float64, n)
	}
	return retArr
}

//Multiply C = AB
//where c is the return matrix
func multiply(a, b Matrix) Matrix {
	//check dimensions of matrix
	if a.n != b.m {
		panic(fmt.Sprintf("Dimensions mismatch a: %v x %v and b: %v x %v",
			a.m, a.n, b.m, b.n))
	}
	retArr := create2DSlice(a.m, b.n)
	for i := 0; i < a.m; i++ {
		for j := 0; j < b.m; j++ {
			sum := 0.0
			for k := 0; k < b.n; k++ {
				sum += a.matrix[i][k] * b.matrix[k][j]
			}
			retArr[i][j] = sum
		}
	}

	return Matrix{retArr, a.n, b.n}
}

//Returns true if two matricies are equal
func equal(a, b Matrix) bool {
	if a.m != b.m || a.n != b.n {
		return false
	}
	for i := 0; i < a.m; i++ {
		for j := 0; j < a.n; j++ {
			if a.matrix[i][j] != b.matrix[i][j] {
				return false
			}
		}
	}
	return true
}

//simple naive matrix transpose method
//dealing with small matricies for now
func (m Matrix) T() Matrix {
	retArr := create2DSlice(m.n, m.m)
	for i := 0; i < m.m; i++ {
		for j := 0; j < m.n; j++ {
			retArr[j][i] = m.matrix[i][j]
		}
	}
	return Matrix{retArr, m.n, m.m}
}

//returns a deep copy of the matrix
func deep_copy(m Matrix) Matrix {
	retArr := create2DSlice(m.m, m.n)
	for i := 0; i < m.m; i++ {
		for j := 0; j < m.n; j++ {
			retArr[i][j] = m.matrix[i][j]
		}
	}
	return Matrix{retArr, m.m, m.n}
}

//QR decomposition of a matrix
//via householder matricies
//Note this overwrites original matrix
//returns 2 matricies
func QR_old(inMat Matrix) (Matrix, Matrix) {
	//create 2 return matricies
	retArr := deep_copy(inMat)
	//r := create2DSlice(inMat.m, inMat.n)
	//Iterater over each column
	for j := 0; j < retArr.n; j++ {
		//get column  normalization factor
		colNorm := 0.0
		for idx := 0; idx < retArr.m; idx++ {
			colNorm += math.Pow(retArr.matrix[j][idx], 2)
		}
		//if norm is 0 then matrix is linearly dependent
		if colNorm == 0 {
			panic(fmt.Sprintf("Matrix has linearly dependent columns"))
		}
		colNorm = 1 / math.Sqrt(colNorm)
		print(colNorm, "\n")
		//norm column
		for i := 0; i < retArr.m; i++ {
			retArr.matrix[j][i] *= colNorm
		}
	}
	return I(2), I(2)
}

// Normalize values in array
// returns scalar normalization of vector
func vecNorm(v []float64) float64 {
	return 1 / sumSq(v)
}

// Sum sq of values in an array
// returns sumsq of values in an array
func sumSq(v []float64) float64 {
	sumsq := 0.0
	for _, num := range v {
		sumsq += math.Pow(num, 2)
	}
	return math.Sqrt(sumsq)
}

// add two values in an array
// returns scalar normalization of vector
func vecAdd(v1, v2 []float64) []float64 {
	if len(v1) != len(v2) {
		panic(fmt.Sprintf("Cannot add two uneven vectors v1=%v, v2=%v", len(v1), len(v2)))
	}
	ret := make([]float64, len(v1))
	for idx, num := range v1 {
		ret[idx] = num + v2[idx]
	}
	return ret
}

//QR decomposition of a matrix
//via householder matricies
//Note this overwrites original matrix
//returns 2 matricies
func QR(inMat Matrix) (Matrix, Matrix) {
	//create 2 return matricies
	//retArr := deep_copy(inMat)
	//r := create2DSlice(inMat.m, inMat.n)
	//Iterater over each column
	for k := 0; k < inMat.n; k++ {
		//get y
		y := make([]float64, inMat.m-k)
		for j := k; j < inMat.m; j++ {
			y[j-k] = inMat.matrix[j][k]
		}
		//w = y + sign(y1)||y||
		y_sumSq := sumSq(y)
		sign := y[0] / math.Abs(y[0])
		for idx, ele := range y {
			y[idx] = ele + sign*y_sumSq
		}
		//get w whose var name is y
		//reusing same variable
		//v = 1/||w||
		y_sumSq = sumSq(y)
		for idx, ele := range y {
			y[idx] = ele * y_sumSq
		}
		fmt.Println(y, sign)
	}
	return I(2), I(2)
}

func main() {
	x := genRandMatrix(3, 3)
	//y := I(5)
	//z := multiply(x, y)

	fmt.Println(x)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(QR(x))
	//fmt.Println(y)
}
