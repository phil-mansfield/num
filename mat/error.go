package mat

import (
	"fmt"
	"runtime"
)

// MatrixError is error type returned by functions which do not result in
// a new Matrix. *MatrixError implements the error interface.
type MatrixError struct{
	Code int // error code representing the type of error which occured
	OperationName string // name of operation which gave error
	Description string // description of especifics of error
	Stack string // stacktrace
}

const (
	// Error codes that can be returned by operations in package mat.
	_ int = iota
	NilError // An operation was called on a nil pointer.
	ShapeError // Requirements on Matrix shapes were not met.
	SingularError // Matrix was singular in a context where this was not
	              // allowed.
	ParameterError // A non-Matrix function parameter was outside the
	               // aceptable range.

	defaultStackSize = 1 << 9 
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
// By default, operations will not panic. Changing this behavior outside of a
// debugging context is strongly discouraged.
func TogglePanic() bool {
	willPanic = !willPanic
	return willPanic
}

// IsError indicates whether m is an error Matrix. IsError returns true if m
// is the result of an invalid operation or if one of the matrices used as
// arguments to this operation was an error Matrix. If m was an error Matrix
// prior to being the target of an operation and the operation succeeds, it
// will no longer be an error Matrix.
func (m *Matrix) IsError() bool {
	return m.err != nil
}

// Error returns a string representing the first error that occured in the
// creation of m. If no errors occured or if m is nil, an empty string is
// returned.
func (m *Matrix) Error() string {
	return m.err.Error()
}

// MatrixError returns a pointer to the struct representing the first error
// that occured in the creation of m. If no such error occured, or if m is nil,
// nil is returned.
func (m *Matrix) MatrixError() *MatrixError {
	return m.err
}

// Error returns a string representation of err. If me is nil, an empty string
// is returned.
func (err *MatrixError) Error() string {
	if err == nil {
		return ""
	}

	name := codeName(err.Code)

	return fmt.Sprintf("%s: %s - mat.%s", name,
		err.Description, err.OperationName)
}

// codeName returns a string representation of the given error code.
func codeName(code int) string {
	switch code {
	case NilError:
		return "Nil Error"
	case ShapeError:
		return "Shape Error"
	case SingularError:
		return "Singular Error"
	case ParameterError:
		return "Parameter Error"
	default:
		panic(fmt.Sprintf("Internal Error: Unrecognized error code: %d", code))
	}
}

// newErrorMatrix creates a new error Matrix corresponding to the given error
// code. operationName should given the name of the function which this function
// is being called from (this will not neccesarily be the name seen by the
// user), and a brief but meaningful description of the error, even if it is
// redundant with the error code. The description should contain no information
// about the operation that generated it, so that the operation may be changed
// as the error is propogated up the stack. All descriptions must be full
// sentences.
func newErrorMatrix(code int, operationName, desc string) *Matrix {
	return &Matrix{[]float64{}, 0, 0, newError(code, operationName, desc)}
}

// newError creates a new MatrixError corresponding to the given error code.
// operationName should given the name of the function which this function is
// being called from (this will not neccesarily be the name seen by the user),
// and a brief but meaningful description of the error, even if it is redundant
// with the error code. The description should contain no information about the
// operation that generated it, so that the operation may be changed as the
// error is propogated up the stack. All descriptions must be full sentences.
func newError(code int, operationName, desc string) *MatrixError {
	err := &MatrixError{code, operationName, desc, ""}
	if willPanic {
		panic(err.Error())
	}

	bytesRead, stackSize := 0, defaultStackSize
	var stackBuf []byte
	for stackSize != bytesRead {
		stackBuf := make([]byte, stackSize)
		bytesRead = runtime.Stack(stackBuf, false)
		stackSize = stackSize << 1
	}
	err.Stack = string(stackBuf[:bytesRead])
	
	return err
}