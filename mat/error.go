package mat

// MatrixError is error type returned by functions which do not result in
// a new Matrix. *MatrixError implements the error interface.
type MatrixError struct{
	Code int // error code representing the type of error which occured
	Description string // description of especifics of error
	OperationName string // name of operation which gave error
	Stack string // stacktrace
}

const (
	// Error codes that can be returned by operations in package mat.
	_ int = iota
	NilError // An operation was called on a nil pointer.
	ShapeError // Requirements on Matrix shapes were not met.
	ParameterError // A non-Matrix function parameter was outside the
	               // aceptable range.
)

var (
	willPanic bool = false
)

// TogglePanic changes the behavior of all functions in package mat when they
// encounter an error. If they currently return error structs, they will
// instead panic with the Error string of that struct as a parameter and
// visa-versa. If the behavior after the funciton has been called is to panic,
// true is returned, otherwise false is returned.
//
// By default operations will not panic.
func TogglePanic() bool {
	return false
}

// Error returns a string representing the first error that occured in the
// creation of m. If no errors occured or if m is nil, an empty string is
// returned.
func (m *Matrix) Error() string {
	return ""
}

// MatrixError returns a pointer to the struct representing the first error
// that occured in the creation of m. If no such error occured, or if m is nil,
// nil is returned.
func (m *Matrix) MatrixError() *MatrixError {
	return nil
}

// Error returns a string representation of err. If me is nil, nil is returned.
func (err *MatrixError) Error() string {
	return ""
}