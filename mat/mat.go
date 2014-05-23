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