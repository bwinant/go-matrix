package matrix

import (
	"fmt"
	"errors"
	"math"
	"bytes"
)

type Matrix interface {
	Rows() int
	Cols() int
	Row(i int) []float64
	Col(j int) []float64
	Get(i, j int) float64
	Set(i int, j int, v float64)
	Max() float64
	Min() float64
	Add(m Matrix) (Matrix, error)
	Sub(m Matrix) (Matrix, error)
	Scalar(x float64) Matrix
	Multiply(m Matrix) (Matrix, error)
	Transpose() Matrix
}

type DenseMatrix struct {
	rows int
	cols int
	data [] float64
}

func NewMatrix(rows, cols int) Matrix {
	if rows <= 0 || cols <= 0 {
		panic(fmt.Sprintf("Invalid matrix dimension: %v x %v", rows, cols))
	}
	return &DenseMatrix{rows: rows, cols: cols, data: make([]float64, rows * cols) }
}

func InitMatrix(n [][] float64) Matrix {
	rows := len(n)
	cols := len(n[0])

	m := &DenseMatrix{rows: rows, cols: cols, data: make([]float64, rows * cols) }

	for i := 1; i <= rows; i++ {
		if len(n[i - 1]) != cols {
			panic(fmt.Sprintf("Invalid row length %v, expected %v", len(n[i - 1]), cols))
		}

		for j := 1; j <= cols; j++ {
			idx := m.index(i, j)
			m.data[idx] = n[i - 1][j - 1]
		}
	}

	return m
}

func IdentityMatrix(dimension int) Matrix {
	if dimension <= 0 {
		panic(fmt.Sprintf("Invalid identity matrix dimension: %v", dimension))
	}

	m := &DenseMatrix{rows: dimension, cols: dimension, data: make([]float64, dimension * dimension) }
	for i := 1; i <= dimension; i++ {
		m.Set(i, i, 1)
	}
	return m
}

func (m *DenseMatrix) Rows() int {
	return m.rows
}

func (m *DenseMatrix) Cols() int {
	return m.cols
}

func (m *DenseMatrix) Row(i int) []float64 {
	r := make([]float64, m.cols)
	start := (i - 1) * m.cols
	end := start + m.cols
	copy(r[:], m.data[start:end])
	return r
}

func (m *DenseMatrix) Col(j int) []float64 {
	c := make([]float64, m.rows)
	idx := j - 1
	for i:= 0; i < m.rows; i++ {
		c[i] = m.data[idx]
		idx += m.cols
	}
	return c
}

func (m *DenseMatrix) Get(i, j int) float64 {
	idx := m.index(i, j)
	return m.data[idx]
}

func (m *DenseMatrix) Set(i int, j int, v float64) {
	idx := m.index(i, j)
	m.data[idx] = v
}

func (m *DenseMatrix) Max() float64 {
	max := 0.0
	for _, v := range m.data {
		if v > max {
			max = v
		}
	}
	return max
}

func (m *DenseMatrix) Min() float64 {
	min := math.MaxFloat64
	for _, v := range m.data {
		if v < min {
			min = v
		}
	}
	return min
}

func (m *DenseMatrix) Add(n Matrix) (Matrix, error) {
	if m.rows != n.Rows() || m.cols != n.Cols() {
		return nil, errors.New("Matrix dimensions do not match")
	}

	r := m.clone()
	for i := 1; i <= r.rows; i++ {
		for j := 1; j <= r.cols; j++ {
			idx := r.index(i, j)
			r.data[idx] = r.data[idx] + n.Get(i, j)
		}
	}
	return r, nil
}

func (m *DenseMatrix) Sub(n Matrix) (Matrix, error) {
	if m.rows != n.Rows() && m.cols != n.Cols() {
		return nil, errors.New("Matrix dimensions do not match")
	}

	r := m.clone()
	for i := 1; i <= r.rows; i++ {
		for j := 1; j <= r.cols; j++ {
			idx := r.index(i, j)
			r.data[idx] = r.data[idx] - n.Get(i, j)
		}
	}
	return r, nil
}

func (m *DenseMatrix) Scalar(x float64) Matrix {
	r := m.clone()
	for i := 1; i <= r.rows; i++ {
		for j := 1; j <= r.cols; j++ {
			idx := r.index(i, j)
			r.data[idx] = r.data[idx] * x
		}
	}
	return r
}

func (m *DenseMatrix) Multiply(n Matrix) (Matrix, error) {
	if m.cols != n.Rows() {
		return nil, errors.New("Matrix dimensions do not match")
	}

	r := &DenseMatrix{rows: m.rows, cols: n.Cols(), data:make([]float64, m.rows * n.Cols()) }
	for i := 1; i <= r.rows; i++ {
		for j := 1; j <= r.cols; j++ {
			sum := 0.0
			for c := 1; c <= m.Cols(); c++ {
				sum += m.Get(i, c) * n.Get(c, j)
			}
			idx := r.index(i, j)
			r.data[idx] = sum
		}
	}

	return r, nil
}

func (m *DenseMatrix) Transpose() Matrix {
	t := &DenseMatrix{rows: m.cols, cols: m.rows, data:make([]float64, m.rows * m.cols) }
	for i := 1; i <= t.rows; i++ {
		for j := 1; j <= t.cols; j++ {
			idx := t.index(i, j)
			t.data[idx] = m.Get(j, i)
		}
	}
	return t
}

func (m *DenseMatrix) String() string {
	max := m.Max()
	digits := 1
	for max > 1 {
		max = max / 10
		digits++
	}
	format := fmt.Sprintf("%%%v.2f ", digits + 3)

	b := new(bytes.Buffer)
	for i := 1; i <= m.rows; i++ {
		b.WriteString("[ ")
		for j := 1; j <= m.cols; j++ {
			b.WriteString(fmt.Sprintf(format, m.Get(i, j)))
		}
		b.WriteString(" ]\n")
	}
	return b.String()
}

func (m *DenseMatrix) index(i, j int) int {
	if i <= 0 || j <= 0 || i > m.rows || j > m.cols {
		panic(fmt.Sprintf("Invalid matrix index: (%v, %v)", i, j))
	}
	return (i - 1) * m.cols + (j - 1)
}

func (m *DenseMatrix) clone() *DenseMatrix {
	r := &DenseMatrix{}
	*r = *m
	return r
}

/*

type Vector struct {
	DenseMatrix
}

func NewVector (rows int) *Vector {
	if rows <= 0 {
		panic(fmt.Sprintf("Invalid vector size: %v", rows))
	}
	return &Vector{DenseMatrix{rows: rows, cols: 1, data:make([]float64, rows)}}
}

func (v *Vector) Get(i, j int) float64 {
	idx := v.index(i)
	return v.data[idx]
}

func (v *Vector) Set(i int, j int, x float64) {
	idx := v.index(i)
	v.data[idx] = x
}

func (v *Vector) At(i int) float64 {
	return v.Get(i, 0)
}

func (v *Vector) SetAt(i int, x float64) {
	v.Set(i, 0, x)
}

func (v *Vector) index(i int) int {
	if i <= 0 || i > v.rows {
		panic(fmt.Sprintf("Invalid vector index: %v", i))
	}
	return i - 1
}*/
