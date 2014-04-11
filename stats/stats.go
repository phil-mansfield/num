/*
Package stats is a collection of basic statistical tests which can be performed
upon float64 arrays.

All operations will panic if given an array of length 0 or an array containing a NaN.
Some operations have additional panic conditions.
*/
package stats

import (
	"fmt"
	"math"
)

// Description is a tuple which stores the values of various basic statistical tests
// performed on an array. Len is the number of values between LoLim and HiLim (inclusive)
// and Mean, Variance, StdDev, Max, and Min are the statistical properties of the
// array contents which are between these two values.
//
// Specifications about these fields can be found in their corresponding functions.
type Description struct {
	Len int
	Mean, Variance, StdDev, Max, Min float64
	MaxIdx, MinIdx int
	LoLim, HiLim float64
}

// BoundedDescribe computes all the fields in a Description tuple for the given array.
// All values lower than loLim or higher than hiLim are ignored.
//
// BoundedDescribe panics if given an array of length 0, if given an array containing Nan,
// if loLim >= hiLim, if either limit is NaN, or if there are
// no array elements between the two bounds. Unlike Mean, BoundedDescribe does not
// panic if given an array containing +Inf and -Inf, so it is possible for Mean to be
// NaN. This is done so that arrays containing
// these values remain valid targets for other statistical tests
func BoundedDescribe(xs []float64, loLim, hiLim float64) *Description {
	if len(xs) <= 0 {
		panic("stats.TrimDescribe given array of length zero.")
	} else if math.IsNaN(hiLim) || math.IsNaN(loLim) || hiLim <= loLim {
		panic(fmt.Sprintf("stats.TrimDescribe given invalid bounds: [%g, %g].",
			loLim, hiLim))
	}
	
	count := 0
	sum := 0.0
	sqrSum := 0.0
	minIdx := 0
	maxIdx := 0
	max := xs[0]
	min := xs[0]

	for i := 0; i < len(xs); i++ {
		if math.IsNaN(xs[i]) {
			panic(fmt.Sprintf("stats.TrimDescribe encountered NaN at index %d.", i))
		}
		if loLim > xs[i] || hiLim < xs[i] { continue }

		count += 1
		sum += xs[i]
		sqrSum += xs[i] * xs[i]
		if xs[i] > max {
			max = xs[i]
			maxIdx = i
		} else if xs[i] < min {
			min = xs[i]
			minIdx = i
		}
	}

	if count == 0 {
		panic(fmt.Sprintf("stats.TrimDescribe given array with no elements in [%g, %g].",
			loLim, hiLim))
	}

	as := new(Description)
	as.Len = count
	as.Mean = sum / float64(as.Len)
	as.Variance = sqrSum / float64(as.Len) - as.Mean * as.Mean
	as.StdDev = math.Sqrt(as.Variance)
	as.Max = max
	as.MaxIdx = maxIdx
	as.Min = min
	as.MinIdx = minIdx
	as.LoLim = loLim
	as.HiLim = hiLim

	// NaN could only arise from infinite-valued array elements:
	if math.IsNaN(as.Variance) {
		as.Variance = math.Inf(+1)
		as.StdDev = math.Inf(+1)
	}

	return as
}

// Describe computes all the fields in a Description tuple for the given array.
//
// Describe panics if given an array of length 0 or if given an array containing Nan.
// Unlike Mean, Describe does not panic if given an array containing +Inf and
// -Inf, so it is possible for Mean to be  NaN. This is done so that arrays containing
// these values remain valid targets for other statistical tests.
func Describe(xs []float64) *Description {
	if len(xs) <= 0 { panic("stats.Describe given array of length zero.") }
	return BoundedDescribe(xs, math.Inf(-1), math.Inf(+1))
}

// Mean returns the average value of the given array.
//
// Mean panics if given an array of length 0, if given an array containing Nan,
// of if given an array containing both + and - infinity.
func Mean(xs []float64) float64 { 
	if len(xs) <= 0 { panic("stats.Mean given array of length zero.") }
	as := Describe(xs)
	if math.IsInf(as.Max, +1) && math.IsInf(as.Min, -1) {
		panic(fmt.Sprintf("stats.Mean called on array containing +Inf at %d and -Inf at %d.",
			as.MaxIdx, as.MinIdx))
	}
	return as.Mean
}

// Variance returns the square of the average deviation from the mean of a given
// array.
//
// Variance panics if given an array of length 0 or if given an array containing Nan.
func Variance(xs []float64) float64 {
	if len(xs) <= 0 { panic("stats.Variance given array of length zero.") }
	return Describe(xs).Variance
}

// StdDev reutrns the average deviation from the mean of a given array.
//
// Mean panics if given an array of length 0 or if given an array containing Nan.
func StdDev(xs []float64) float64 {
	if len(xs) <= 0 { panic("stats.StdDev given array of length zero.") }
	return Describe(xs).StdDev
}

// Max returns the largest value in the given array as well as the first index
// at which this value occured.
//
// Max panics if given an array of length 0 or if given an array containing Nan.
func Max(xs []float64) (float64, int) {
	if len(xs) <= 0 { panic("stats.Max given array of length zero.") }
	as := Describe(xs)
	return as.Max, as.MaxIdx
}

// Max returns the smallest value in the given array as well as the first index
// at which this value occured.
///
// Min panics if given an array of length 0 or if given an array containing Nan.
func Min(xs []float64) (float64, int) {
	if len(xs) <= 0 { panic("stats.Min given array of length zero.") }
	as := Describe(xs)
	return as.Min, as.MinIdx
}
