/*
package mat implements basic matrix operations on real-valued matrices. The 
subpackage cmat/ implements the operations for complex-valued matrices. In the
interest of providing clean interfaces several non-trivial optimizations based
on parameter-spamming and algorithm selection have been ignored. These
clunkier, optimized interfaces can be found in the optmat/ and optcmat/
subpackages.
*/
package mat

// Matrix represents a two-dimensional rectangluar array of real values.
type Matrix struct { }

// MatrixError is the error type used by package mat. *MatrixError implements
// the error interface.
type MatrixError struct { }

func (err *MatrixError) Error() string {
	return ""
}

// type assertion
var _ error = &MatrixError{}

// Initialization functions

// New returns a matrix with the given dimensions where all elements are
// initialized to zero.
//
// If height or width is non-positive, an error will be returned.
func New(width, height int) *Matrix, error {
	return nil, nil
}

// Identity returns a square matrix with the given width which contains ones
// down its diagonal and zeroes everywhere else.
//
// If height or width is non-positive, an error will be returned.
func Identity(width int) *Matrix, error {
	return nil, nil
}

// FromArray converts an slice of floats to a matrix with the given dimensions.
// The element with zero-indexed coordinates (x, y) in the matrix will be the
// same as that at index data[y * width + x].
//
// If width * height != len(data) or if height or width is non-positive, an
// error will be returned.
func FromArray(width, height int, data []float64) *Matrix, error {
	return nil, nil
}

// FromGrid conversts a 2D slice of floats to a matrix with the same
// dimensions. The element with zero-indexed coordinates of (x, y) in the
// matrix will be the same as that at  index data[y][x].
//
// If any two rows in data have different lengths, or if len(data) == 0 or
// len(data[0]) == 0, an error will be returned.
func FromGrid(data [][]float64) *Matrix, error {
	return nil, nil
}