package mat

// Eigenvalues returns a slice containing the eigenvalues of m. The order
// of the eigenvalues is the same as the order of the corresponding
// eigenvectors returned by m.Eigenvectors().
//
// If the Matrix is not square, an error is returned.
func (m *Matrix) Eigenvalues() ([]complex128, error) {
	return nil, nil
}

// Eigenvectors returns a 2D slice containing the eigenvectors of m.
// The order of the eigenvectos is the same as the order of the corresponding
// eigenvalues returned by m.Eigenvalues().
//
// If the Matrix is not a square matrix, an error is returned.
func (m *Matrix) Eigenvectors() ([][]complex128, error) {
	return nil, nil
}

// Determinant returns the determinant of m.
//
// If m is not a square matrix, an error is returned.
func (m *Matrix) Determinant() (float64, error) {
	return -1, nil
}

// Invert returns the inverse of m.
//
// If m is not a square matrix or is sigular, an error is returned.
func Invert(m *Matrix) (*Matrix, error) {
	return nil, nil
}

// Invert calculates the inverse of m and places it in the target Matrix.
//
// If m is not a square matrix, an error is returned.
func (target *Matrix) Invert(m *Matrix) error {
	return nil
}