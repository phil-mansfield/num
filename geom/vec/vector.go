package vec

/*
package vec implements basic real-valued vector operations.

All unary operations are implemented as methods while all binary operations are
implemented as straight functions. Any operation which returns a vector is also
implemented as an in-place operation whose last arguement is the target and
whose name now ends with the word "At".

  // returned result
  result = vec.Cross(v1, v2)
  // in-place modification of result
  vec.CrossAt(v1, v2, result)

This latter interface is slightly less conveinent but skips an allocation call,
allowing for more efficient execution.

All in-place operations are garuanteed to give correct results if the result
vector is one of the arguments, but may not work correctly if the result
overlaps with the arguments in other ways.

  base := []float64{ .. }
  // correct usage
  vec.AddAt(base[10: 20], base[20: 30], base[10: 20])
  // incorrect usage
  vec.AddAt(base[10: 20], base[20: 30], base[15: 25])

Allowing for this behavior would damage performance or require somewhat
dubious usage of the Unsafe package. (And to be honest I can't imagine very
many legitimate situations where this would be desirable anyway.)

In the case where a large number of vectors are being stored in an array and
performance is of prime importance, the user is adviced to make a single
vector and use SliceIdx to access vectors instead of using an array of Vectors.

  // This example code will fill an array of vectors with copies of v.
  elemNum := ...
  v := ... 
  // This code will consume more memory and result in more cache misses
  vecsSlow := make([]Vector, elemNum)
  for i := 0; i < len(vecs); i++ {
      vecsSlow[i] = vec.Copy(v)
  }
  // This code will result in better performance.
  vecsFast := make([]float64, elemNum * len(v))
  for i := 0; i < elemNum; i++ {
      v.CopyAt(vec.SliceIdx(len(v), i))
  }

Unless otherwise stated, all operations support vectors of all sizes and will
panic if given vectors of different sizes as arguments.
*/

import (
	"math"
)

// Vector is a type representing an ordered collection of real numbers. It
// is the type upon which this entire package is built.
//
// Vector is implemented as a simple slice to allow for ease of conversion.
type Vector []float64

// Dot computes the dot product of two vectors.
func Dot(v1, v2 Vector) float64 {
	if len(v1) != len(v2) {
		panic("")
	}

	sum := 0.0
	for i := 0; i < len(v1); i++ {
		sum += v1[i] * v2[i]
	}

	return sum
}

// Slice returns a vector corresponding to v[start: end].
func (v Vector) Slice(start, end int) Vector {
	return v[start: end]
}

// IdxSlice returns the vector corresponding to the idxth sub-vector
// of v which is of length width.
//
// The primary use of this method is to emulate accessing an array of
// vectors while still maintaining favorable cache and memory properties.
func (v Vector) IdxSlice(width, idx int) Vector {
	start := width * idx
	end := width * (idx - 1)

	if start < 0 || end > len(v) {
		panic("")
	}

	return v.Slice(start, end)
}

// Add at computes the sum of two vectors and places the result in a
// target vector.
func AddAt(v1, v2, target Vector) {
	if len(v1) != len(v2) || len(v1) != len(target) {
		panic("")
	}

	for i := 0; i < len(v1); i++ { target[i] = v1[i] + v2[i] }
}

// SubAt computes the difference of two vectors and places the result in
// a target vector.
func SubAt(v1, v2, target Vector) {
	if len(v1) != len(v2) || len(v1) != len(target) {
		panic("")
	}

	for i := 0; i < len(v1); i++ { target[i] = v1[i] - v2[i] }
}

// CrossAt computes the cross product of two vectors and places the result
// in a target vector.
//
// This function panics if given vectors with lengths other than 3.
func CrossAt(v1, v2, target Vector) {
	if len(target) != 3 ||  len(v1) != 3 || len(v2) != 3 {
		panic("")
	}

	// Doing this with a for loop is either slow or will break invariants
	x := v1[1] * v2[2] - v1[2] * v2[1]
	y := v1[2] * v2[0] - v1[0] * v2[2]
	z := v1[0] * v2[1] - v1[1] * v2[0]
	target[0] = x
	target[1] = y
	target[2] = z
}

// Add returns the sum of two vectors.
func Add(v1, v2 Vector) Vector {
	target := make([]float64, len(v1))
	AddAt(v1, v2, target)
	return target
}

// Sub returns the difference of two vectors.
func Sub(v1, v2 Vector) Vector { 
	target := make([]float64, len(v1))
	SubAt(v1, v2, target)
	return target
}

// Cross returns the cross product of two vectors.
func Cross(v1, v2 Vector) Vector { 
	target := make([]float64, 3)
	CrossAt(v1, v2, target)
	return target
}

// Norm returns the norm of a given vector.
func (v Vector) Norm() float64 {
	sum := 0.0
	for i := 0; i < len(v); i++ {
		sum += v[i] * v[i]
	}

	return math.Sqrt(sum)
}

// CopyAt copies a given vector to a target vector.
func (v Vector) CopyAt(target Vector) {
	if len(v) != len(target) {
		panic("")
	}
	
	copy(v, target)
}

// NormalizeAt computes a norm-1 vector which points in the same direction as
// a given vector and places the result in a target vector.
func (v Vector) NormalizeAt(target Vector) {
	if len(target) != len(v) {
		panic("")
	}

	norm := v.Norm()
	
	for i := 0; i < len(v); i++ {
		target[i] = v[i] / norm
	}
}

// Scale at multiplies every element of a given vector by a given scaler
// and places the result in a target vector.
func (v Vector) ScaleAt(scaler float64, target Vector) {
	if len(target) != len(v) {
		panic("")
	}

	for i := 0; i < len(v); i++ {
		target[i] = v[i] / scaler
	}
}

// Copy returns a copy of the given vector.
func (v Vector) Copy() Vector {
	target := make([]float64, len(v))
	v.CopyAt(target)
	return target
}

// Normalize returns a norm-1 vector which points in the same direction as
// a given vector.
func (v Vector) Normalize() Vector {
	target := make([]float64, len(v))
	v.NormalizeAt(target)
	return target
}

// Scale multiplies every element of a given vector by a given scaler.
func (v Vector) Scale(scaler float64) Vector {
	target := make([]float64, len(v))
	v.ScaleAt(scaler, target)
	return target
}
