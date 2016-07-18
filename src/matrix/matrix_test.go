package matrix

import (
	"testing"
)

func TestInitMatrix(t *testing.T) {
	m := InitMatrix([][]float64 {
		{10, 9, 8},
		{0, 1, 3},
		{4, 5, 6},
		{0, 1, 0.5},
	})

	if m.Rows() != 4 && m.Cols() != 3 {
		t.Fail()
	}

	if m.Get(1, 1) != 10 {
		t.Fail()
	}

	if m.Get(2, 2) != 1 {
		t.Fail()
	}

	if m.Get(4, 3) != 0.5 {
		t.Fail()
	}
}

func TestMatrix_Dim(t *testing.T) {
	var m Matrix

	m = NewMatrix(1, 1)
	if m.Rows() != 1 && m.Cols() != 1 {
		t.Fail()
	}

	m = NewMatrix(2, 3)
	if m.Rows() != 2 && m.Cols() != 3 {
		t.Fail()
	}
}

func TestMatrix_Row(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 0, 8, 5},
		{2, 5, 12, 0.5},
		{3, 1, .25, 7},
		{5, -1, 10, 20},
	})

	r1 := m.Row(1)
	if len(r1) != 4 {
		t.Fail()
	}
	if r1[0] != 1 {
		t.Fail()
	}
	if r1[1] != 0 {
		t.Fail()
	}
	if r1[2] != 8 {
		t.Fail()
	}
	if r1[3] != 5 {
		t.Fail()
	}

	r3 := m.Row(3)
	if len(r3) != 4 {
		t.Fail()
	}
	if r3[0] != 3 {
		t.Fail()
	}
	if r3[1] != 1 {
		t.Fail()
	}
	if r3[2] != .25 {
		t.Fail()
	}
	if r3[3] != 7 {
		t.Fail()
	}

	r4 := m.Row(4)
	if len(r4) != 4 {
		t.Fail()
	}
	if r4[0] != 5 {
		t.Fail()
	}
	if r4[1] != -1 {
		t.Fail()
	}
	if r4[2] != 10 {
		t.Fail()
	}
	if r4[3] != 20 {
		t.Fail()
	}
}

func TestMatrix_Col(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 0, 8, 5, -6},
		{2, 5, 12, 0.5, 1.25},
		{3, 1, .25, 7, -1.5},
		{5, -1, 10, 20, 0},
	})

	c1 := m.Col(1)
	if len(c1) != 4 {
		t.Fail()
	}
	if c1[0] != 1 {
		t.Fail()
	}
	if c1[1] != 2 {
		t.Fail()
	}
	if c1[2] != 3 {
		t.Fail()
	}
	if c1[3] != 5 {
		t.Fail()
	}

	c3 := m.Col(3)
	if len(c3) != 4 {
		t.Fail()
	}
	if c3[0] != 8 {
		t.Fail()
	}
	if c3[1] != 12 {
		t.Fail()
	}
	if c3[2] != .25 {
		t.Fail()
	}
	if c3[3] != 10 {
		t.Fail()
	}

	c5 := m.Col(5)
	if len(c5) != 4 {
		t.Fail()
	}
	if c5[0] != -6 {
		t.Fail()
	}
	if c5[1] != 1.25 {
		t.Fail()
	}
	if c5[2] != -1.5 {
		t.Fail()
	}
	if c5[3] != 0 {
		t.Fail()
	}
}

func TestMatrix_SetGet(t *testing.T) {
	var m Matrix

	m = NewMatrix(1, 1)
	m.Set(1, 1, 5)
	if m.Get(1, 1) != 5 {
		t.Fail()
	}
}

func TestMatrix_Add(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 0},
		{2, 5},
		{3, 1},
	})

	n := InitMatrix([][]float64 {
		{4, 0.5},
		{2, 5},
		{0, 1},
	})

	r, err := m.Add(n)
	if err != nil {
		t.Errorf("Add returned an error: %v", err)
	}

	if r.Get(1, 1) != 5 {
		t.Errorf("Add returned incorrect value %v at (%v, %v)", r.Get(1, 1), 1, 1)
	}
	if r.Get(1, 2) != 0.5 {
		t.Errorf("Add returned incorrect value %v at (%v, %v)", r.Get(1, 2), 1, 2)
	}
	if r.Get(2, 1) != 4 {
		t.Errorf("Add returned incorrect value %v at (%v, %v)", r.Get(2, 1), 2, 1)
	}
	if r.Get(2, 2) != 10 {
		t.Errorf("Add returned incorrect value %v at (%v, %v)", r.Get(2, 2), 2, 2)
	}
	if r.Get(3, 1) != 3 {
		t.Errorf("Add returned incorrect value %v at (%v, %v)", r.Get(3, 1), 3, 1)
	}
	if r.Get(3, 2) != 2 {
		t.Errorf("Add returned incorrect value %v at (%v, %v)", r.Get(3, 2), 3, 2)
	}
}

func TestMatrix_Add_Error(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 2},
		{0, 1},
		{4, 5},
	})

	n := InitMatrix([][]float64 {
		{1, 1},
		{0, 0},
	})

	r, err := m.Add(n)
	if r != nil || err == nil {
		t.Errorf("Add did not return error")
	}
}

func TestMatrix_Scalar(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 2},
		{0, 1},
		{4, 5},
	})

	r := m.Scalar(2)

	if r.Get(1, 1) != 2 {
		t.Fail()
	}
	if r.Get(1, 2) != 4 {
		t.Fail()
	}
	if r.Get(2, 1) != 0 {
		t.Fail()
	}
	if r.Get(2, 2) != 2 {
		t.Fail()
	}
	if r.Get(3, 1) != 8 {
		t.Fail()
	}
	if r.Get(3, 2) != 10 {
		t.Fail()
	}
}

func TestMatrix_Multiply(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 3, 2},
		{4, 0, 1},
	})

	n := InitMatrix([][]float64 {
		{1, 3},
		{0, 1},
		{5, 2},
	})

	r, err := m.Multiply(n)
	if err != nil {
		t.Errorf("Multiply returned an error: %v", err)
	}

	if r.Rows() != 2 && r.Cols() != 2 {
		t.Fail()
	}

	if r.Get(1, 1) != 11 {
		t.Fail()
	}
	if r.Get(1, 2) != 10 {
		t.Fail()
	}
	if r.Get(2, 1) != 9 {
		t.Fail()
	}
	if r.Get(2, 2) != 14 {
		t.Fail()
	}
}

func TestMatrix_Multiply2(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 3},
		{2, 5},
	})

	n := InitMatrix([][]float64 {
		{0, 1},
		{3, 2},
	})

	r, err := m.Multiply(n)
	if err != nil {
		t.Errorf("Multiply returned an error: %v", err)
	}

	if r.Rows() != 2 && r.Cols() != 2 {
		t.Fail()
	}

	if r.Get(1, 1) != 9 {
		t.Fail()
	}
	if r.Get(1, 2) != 7 {
		t.Fail()
	}
	if r.Get(2, 1) != 15 {
		t.Fail()
	}
	if r.Get(2, 2) != 12 {
		t.Fail()
	}
}

func TestMatrix_Multiply_Error(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 3, 2, 9},
		{4, 0, 1, -1},
	})

	n := InitMatrix([][]float64 {
		{1, 3},
		{0, 1},
		{5, 2},
	})

	r, err := m.Multiply(n)
	if r != nil || err == nil {
		t.Errorf("Multiply did not return error")
	}
}

func TestMatrix_Max(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 3.7},
		{-2, 12},
		{50, 200},
	})

	if m.Max() != 200 {
		t.Fail()
	}
}

func TestMatrix_Min(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 3.7},
		{-2, 1},
		{5, 2},
	})

	if m.Min() != -2 {
		t.Fail()
	}
}

func TestMatrix_Transpose(t *testing.T) {
	m := InitMatrix([][]float64 {
		{1, 2, 0},
		{3, 5, 9},
	})

	r := m.Transpose()

	if r.Rows() != 3 || r.Cols() != 2 {
		t.Fail()
	}

	if r.Get(1, 1) != 1 {
		t.Fail()
	}
	if r.Get(1, 2) != 3 {
		t.Fail()
	}
	if r.Get(2, 1) != 2 {
		t.Fail()
	}
	if r.Get(2, 2) != 5 {
		t.Fail()
	}
	if r.Get(3, 1) != 0 {
		t.Fail()
	}
	if r.Get(3, 2) != 9 {
		t.Fail()
	}
}

func TestMatrix_NewIdentityMatrix(t *testing.T) {
	i := IdentityMatrix(3)

	for r := 1; r <= i.Rows(); r++ {
		for c := 1; c <= i.Cols(); c++ {
			if r != c && i.Get(r, c) != 0 {
				t.Fail()
			}
			if r == c && i.Get(r, c) != 1 {
				t.Fail()
			}
		}
	}
}
