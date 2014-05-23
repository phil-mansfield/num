package mat

import (
	"github.com/phil-mansfield/num"
)

// Exp computes the matrix exponential of m and returns the result.
//
// If m is nil or not a square Matrix, an error Matrix is returned.
func Exp(m *Matrix) *Matrix {
	return nil
}

// Sin computes the matrix radian sine of m and returns the result.
//
// If m is nil or not a square Matrix, an error Matrix is returned.
func Sin(m *Matrix) *Matrix {
	return nil
}

// Sinh computes the matrix hyperbolic sine of m and returns the result.
//
// If m is nil or not a square Matrix, an error Matrix is returned.
func Sinh(m *Matrix) *Matrix {
	return nil
}

// Cos computes the matrix radian cosine of m and returns the result.
//
// If m is nil or not a square Matrix, an error Matrix is returned.
func Cos(m *Matrix) *Matrix {
	return nil
}

// Cosh computes the matrix hyperbolic cosine of m and returns the result.
//
// If m is nil or not a square Matrix, an error Matrix is returned.
func Cosh(m * Matrix) *Matrix {
	return nil
}

// Log computes the matrix natural logarithm of m and returns the result.
//
// If m is nil or not a square Matrix, an error Matrix is returned.
func Log(m *Matrix) *Matrix {
	return nil
}

// Sqrt computes the matrix square root of m and returns the result.
//
// If m is nil or not a square Matrix, an error Matrix is returned.
func Sqrt(m *Matrix) *Matrix {
	return nil
}

// Func computes the matrix function f(m) and returns the result.
//
// If m is nil or not a square Matrix, an error Matrix is returned.
func Func(f num.Func1D, m *Matrix) *Matrix {
	return nil
}

// Exp computes the matrix exponential of m and stores the result in the
// target Matrix.
//
// If m is nil or not a square Matrix or if target is not the same shape as m,
// target is set to an error Matrix.
func (target *Matrix) Exp(m *Matrix) {
	return
}

// Sin computes the matrix radian sine of m and stores the result in the
// target Matrix.
//
// If m is nil or not a square Matrix or if target is not the same shape as m,
// target is set to an error Matrix.
func (target *Matrix) Sin(m *Matrix) {
	return
}

// Sinh computes the matrix hyperbolic sine of m and stores the result in the
// target Matrix.
//
// If m is nil or not a square Matrix or if target is not the same shape as m,
// target is set to an error Matrix.
func (target *Matrix) Sinh(m *Matrix) {
	return
}

// Cos computes the matrix radian cosine of m and stores the result in the
// target Matrix.
//
// If m is nil or not a square Matrix or if target is not the same shape as m,
// target is set to an error Matrix.
func (target *Matrix) Cos(m *Matrix) {
	return
}

// Cosh computes the matrix hyperbolic cosine of m and stores the result in the
// target Matrix.
//
// If m is nil or not a square Matrix or if target is not the same shape as m,
// target is set to an error Matrix.
func (target *Matrix) Cosh(m * Matrix) {
	return
}

// Log computes the matrix natural logarithm of m and stores the result in the
// target Matrix.
//
// If m is nil or not a square Matrix or if target is not the same shape as m,
// target is set to an error Matrix.
func (target *Matrix) Log(m *Matrix) {
	return
}

// Sqrt computes the matrix square root of m and stores the result in the
// target Matrix.
//
// If m is nil or not a square Matrix or if target is not the same shape as m,
// target is set to an error Matrix.
func (target *Matrix) Sqrt(m *Matrix) {
	return
}

// Func computes the matrix function f(m) and stores the result in the target
// Matrix.
//
// If m is nil or not a square Matrix or if target is not the same shape as m,
// target is set to an error Matrix.
func (target *Matrix) Func(f num.Func1D, m *Matrix) {
	return
}