package mat

// Matrix represents a two-dimensional rectangluar array of real values.
// *Matrix implements the error interface.
type Matrix struct {
	values []float64
	width, height int
	err *MatrixError
}

// New returns a matrix with the given dimensions where all elements are
// initialized to zero.
//
// If height or width is non-positive, an error Matrix will be returned.
func New(width, height int) *Matrix {
	m := new(Matrix)
	m.width, m.height = width, height
	m.values = make([]float64, width * height)

	return m
}

// Identity returns a square matrix with the given width which contains ones
// down its diagonal and zeroes everywhere else.
//
// If height or width is non-positive, an error Matrix will be returned.
func Identity(width int) *Matrix {
	return nil
}

// FromArray converts an slice of floats to a matrix with the given dimensions.
// The element with zero-indexed coordinates (x, y) in the matrix will be the
// same as that at index values[y * width + x].
//
// If width * height != len(values) or if height or width is non-positive, an
// error Matrix will be returned.
func FromSlice(width, height int, values []float64) *Matrix {
	return nil
}

// FromGrid conversts a 2D slice of floats to a matrix with the same
// dimensions. The element with zero-indexed coordinates of (x, y) in the
// matrix will be the same as that at index values[y][x].
//
// If any two rows in data have different lengths, or if len(values) == 0 or
// len(values[0]) == 0, an error Matrix will be returned.
func FromGrid(values [][]float64) *Matrix {
	return nil
}

// IsError indicates whether m is an error Matrix. IsError returns true if m
// is the result of an invalid operation or if one of the matrices used as
// arguments to this operation was an error Matrix. If m was an error Matrix
// prior to being the target of an operation and the operation succeeds, it
// will no longer be an error Matrix.
func (m *Matrix) IsError() bool {
	return false
}

// Equal returns true if every element in the two given arrays is equal to
// within the library precision fraction, ConvergenceEpsilon, as defined in
// num/config.go.
func Equal(m1, m2 *Matrix) bool {
	return false
}

// Compatible returns true if the two given matrices have the same shapes and
// false otherwise. If either Matrix is nil, Compatible returns false.
func Compatible(m1, m2 *Matrix) bool {
	return false
}

// MultCompatible returns true if the two given matrices can be multiplied
// together and false otherwise. If either Matrix is nil, MultCompatible
// returns false.
func MultCompatible(m1, m2 *Matrix) bool {
	return false
}

// TransposeCompatible returns true if m1 is the same shape as the transpose
// of m2 and false otherwise. If either Matrix is nil, TransposeCompatible
// returns false.
func TransposeCompatible(m1, m2 *Matrix) bool {
	return false
}

// Height returns the height of the matrix.
func (m *Matrix) Height() int {
	return -1
}

// Width returns the width of the matrix.
func (m *Matrix) Width() int {
	return -1
}

// Slice returns a slice containing all the values within m. The value at the
// zero-indexed coordinates (x, y) will be placed at index x + m.Width() * y
// in the slice.
func (m *Matrix) Slice() []float64 {
	return nil
}

// Grid returns a 2D slice containing all the values within m. The value at
// the zero-indexed coordinates (x, y) will be placed at index grid[y][x] in
// the output grid.
func (m *Matrix) Grid() [][]float64 {
	return nil
}

// InBounds returns true if the (x, y) coordinate pair is within the bounds
// of m and false otherwise.
func (m *Matrix) InBounds(x, y int) bool {
	return false
}

// Get returns the element of the matrix with coordinates (x, y).
//
// Get and Set are unique in that they panic upon out of bounds input instead
// of returning an error.
func (m *Matrix) Get(x, y int) float64 {
	return 0.0
}

// Set changes the element in the matrix with coordinates (x, y) so that it
// has the given value.
//
// Get and Set are unique in that they panic upon out of bounds input instead
// of returning an error.
func (m *Matrix) Set(x, y int, val float64) {
	return
}

// Print prints the contents of the matrix as a comma-separated array of
// arrays. Each row in the matrix is given its own line.
func (m *Matrix) Print() {
	return
}

// Printf prints the contents of of the matrix as a comma-separated array of
// arrays where each element is formatted according to the given format string.
// Each row in the matrix is given its own line.
func (m *Matrix) Printf(fmt string) {
	return
}

// Copy returns a copy of m.
//
// If m is nil, Copy returns an error Matrix.
func Copy(m *Matrix) *Matrix {
	return nil
}

// Copy copies the values in m to target. The target matrix is also returned.
//
// If target is not the same shape as m or if m is nil, target is set to an
// error Matrix.
func (target *Matrix) Copy(m *Matrix) *Matrix {
	return nil
}
