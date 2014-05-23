package mat

// Add computes m1 + m2 and returns the result.
//
// If m1 and m2 are not the same shape, an error is returned.
func Add(m1, m2 *Matrix) (*Matrix, error) {
	return nil, nil
}

// Sub computes m1 - m2 and returns the result.
//
// If m1 and m2 are not the same size, an error is returned.
func Sub(m1, m2 *Matrix) (*Matrix, error) {
	return nil, nil
}

// Mult computes m1 * m2 and returns the result.
//
// If if the width of m1 is not the same as the height of m2, an error is
// returned.
func Mult(m1, m2 *Matrix) (*Matrix, error) {
	return nil, nil
}

// Scale multiplies every element in m by c and returns the result.
func Scale(m *Matrix, c float64) *Matrix {
	return nil
}

// Add computes m1 + m2 and stores the result in the target Matrix.
//
// If m1, m2, and target are not all the same size, an error is retunred.
func (target *Matrix) Add(m1, m2 *Matrix) error {
	return nil
}

// Sub computes m1 - m2 and stores the result in the target Matrix.
//
// If m1, m2, and target are not all the same size, an error is returned.
func (target *Matrix) Sub(m1, m2 *Matrix) error {
	return nil
}

// Mult computes m1 * m2 and stores the result in the target matrix.
//
// If the width of m1 is not the same as the height or m2, or if target does
// not have the same width as m2 and the same height as m1, an error is
// returned.
func (target *Matrix) Mult(m1, m2 *Matrix) error {
	return nil
}

// Scale multiplies every element of m by c and stores the result in the
// target Matrix.
//
// If m and target are not the same size, an error is returned.
func (target *Matrix) Scale(m *Matrix, c float64) error {
	return nil
}