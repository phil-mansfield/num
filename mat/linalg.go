package mat

// Eigenvalues returns a slice containing the eigenvalues of m. The order
// of the eigenvalues is the same as the order of the corresponding
// eigenvectors returned by m.Eigenvectors().
//
// If m is nil or not square, a non-nil error is returned.
func (m *Matrix) Eigenvalues() ([]complex128, error) {
	return nil, nil
}

// Eigenvectors returns a 2D slice containing the eigenvectors of m.
// The order of the eigenvectos is the same as the order of the corresponding
// eigenvalues returned by m.Eigenvalues().
//
// If m is nil or not a square matrix, a non-nil error is returned.
func (m *Matrix) Eigenvectors() ([][]complex128, error) {
	return nil, nil
}

// Determinant returns the determinant of m.
//
// If m is nil or not a square Matrix, a non-nil error is returned.
func (m *Matrix) Determinant() (float64, error) {
	return 0, nil
}

// Invert returns the inverse of m.
//
// If m is nil or not a square matrix or is singular, an error Matrix is
// returned.
func Invert(m *Matrix) *Matrix {
	return nil
}

// Invert calculates the inverse of m and places it in the target Matrix.
//
// If m is nil or not a square Matrix or is sigular, target is set to an error
// Matrix. target is also set to an error Matrix if it is not the same shape as
// the transpose of m.
func (target *Matrix) Invert(m *Matrix) {
	return
}

// Transpose returns the transpose of m.
//
// If m is nil, an error Matrix is returned.
func Transpose(m *Matrix) *Matrix {
	return nil
}

// Transpose computes the transpose of m and places it in the target Matrix.
//
// If m is nil or if target is not hte same shape as the transpose of m, target
// is set to an error Matrix.
func (target *Matrix) Transpose(m *Matrix) {
	return
}
