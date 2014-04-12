package stats

import (
	"fmt"
	"math"
)

// Histogram is a type which stores the frequency with which float64 values
// fall within finite-width bins.
//
// Histogram may be uniformly distributed in either logarithmic or linear
// space.
type Histogram struct {
	// The number of items in each bin.
	Bins []int
	// The value associated with each bin.
	BinValues []float64
	// BinEdges[i] is the inlcusive lower limit of Bin[i] and the exclusive
	// upper limit of Bin[i - 1]. The final upper limit is inclusive.
	BinEdges []float64
	// The number of values stored in the Histogram.
	ValueCount int

	binWidth float64
	lowLim, highLim float64
	logHistogram bool
}

// Add adds a single value to the Histogram. If the value is out of bounds it
// is ignored and 1 is returned. Otherwise 0 is returned.
func (hist *Histogram) Add(x float64) int {
	var idx int

	if math.IsNaN(x) { panic("stats.Histogram.Add() given NaN.") }
	if hist.logHistogram {
		if x <= 0 { return 1 }
		x = math.Log(x)
	}

	if x == hist.highLim {
		idx = len(hist.Bins) - 1
	} else {
		idx = int((x - hist.lowLim) / hist.binWidth)
		if idx < 0 || idx >= len(hist.Bins) { return 1 }
	}

	hist.ValueCount++
	hist.Bins[idx]++
	return 0
}

// AddArray adds every element of a given array to the Histogram. Elements out
// of bounds are ignored. The return value is the number of such ignored
// elements.
func (hist *Histogram) AddArray(xs []float64) int {
	var idx int
	initialValueCount := 0

	for i, x := range xs {
		if math.IsNaN(x) { 
			panic(fmt.Sprintf("stats.Histogram.AddArray() given NaN at index %d", i))
		}
		if hist.logHistogram {
			if x <= 0 { continue }
			x = math.Log(x)
		}

		if x == hist.highLim {
			idx = len(hist.Bins) - 1
		} else {
			idx = int((x - hist.lowLim) / hist.binWidth)
			if idx < 0 || idx >= len(hist.Bins) { continue }
		}

		hist.ValueCount++
		hist.Bins[idx]++
	}
	
	return hist.ValueCount - initialValueCount + len(xs)
}

// NormalizedBins returns an array with the same relative frequencies as
// hist.Bins, but whose total sum is equal to a given area.
func (hist *Histogram) NormalizedBins(area float64) []float64 {
	valueArea := area / float64(hist.ValueCount)
	
	normedBins := make([]float64, len(hist.Bins))
	for i := 0; i < len(hist.Bins); i++ {
		normedBins[i] = float64(hist.Bins[i]) * valueArea
	}

	return normedBins
}

// CumulativeBins returns an array whose values represent the number of values
// in hist which are in a bin at least as large as the current one.
func (hist *Histogram) CumulativeBins() []int {
	cumBins := make([]int, len(hist.Bins))
	binNum := len(hist.Bins)
	
	cumBins[binNum - 1] = hist.Bins[binNum - 1]
	for i := binNum - 2; i >= 0; i-- {
		cumBins[i] = cumBins[i + 1] + hist.Bins[i]
	}

	return cumBins
}

// NewHistogram creates a Histogram instance out of the given array of values
// with the given number of bins. The minimum and maximum allowed values are
// taken to be the minimum and maximum values in that array.
//
// NewHistogram will panic if given a non-positive number of bins or less than
// two starting values, as either option leads to an ill-defined range.
// NewHistogram will panic fi given any infinite values for the same reason.
func NewHistogram(xs []float64, binNum int) *Histogram {
	if len(xs) <= 1 {
		panic("stats.NewHistogram given empty array.")
	} else if binNum < 1 {
		panic(fmt.Sprintf("stats.NewHistogram given binNum of %d", binNum))
	}

	as := Describe(xs)
	min, max := as.Min, as.Max

	if math.IsInf(min, 0) {
		panic(fmt.Sprintf("stats.NewHistogram given array with infinite value at %d", as.MinIdx))
	} else if math.IsInf(max, 0)  {
		panic(fmt.Sprintf("stats.NewHistogram given array with infinite value at %d", as.MaxIdx))
	}

	hist, _ := NewBoundedHistogram(xs, binNum, min, max)
	return hist
}

// NewHistogram creates a Histogram instance out of the given array of values
// with the given number of bins. Logarithms of bin centers are uniformly
// distributed. The minimum and maximum allowed values are taken to be the
// minimum and maximum values in that array.
//
// NewLogHistogram will panic if given a non-positive number of bins or less than
// two starting values, as either option leads to an ill-defined range. NewLogHistogram
// will also panic if given any negative values or infinite values for the same reason.
func NewLogHistogram(xs []float64, binNum int) *Histogram {
	if len(xs) <= 1 {
		panic("stats.NewLogHistogram given empty array.")
	} else if binNum < 1 {
		panic(fmt.Sprintf("stats.NewLogHistogram given binNum of %d", binNum))
	}

	as := Describe(xs)
	min, max := as.Min, as.Max
	
	if min <= 0 {
		panic(fmt.Sprintf("stats.NewLogHistogram given non-positive value %d at %d", min, as.MinIdx))
	}

	hist, _ := NewBoundedLogHistogram(xs, binNum, min, max)
	return hist
}

// NewBoundedHistogram creates a Histogram instance out of the given array
// of values with the given number of bins which fall between the given
// limits. Any values outside of these limits are ignored. The returned
// integer is the number of such ignored values. Because of this infinite values
// do not cause a panic.
//
// NewBoundedHistogram panics if given a non-positive number of bins or
// a low bound as large or larger than the high bound or if given infinte
// bounds.
func NewBoundedHistogram(xs []float64, binNum int, low, high float64) (*Histogram, int) {
	if binNum < 1 {
		panic(fmt.Sprintf("stats.NewBoundedHistogram given binNum of %d",
			binNum))
	} else if low >= high || math.IsInf(low, 0) || math.IsInf(high, 0) {
		panic(fmt.Sprintf("stats.NewBoundedHistogram given range [%d, %d]",
			low, high))
	}

	hist := new(Histogram)
	hist.Bins = make([]int, binNum)
	hist.BinValues = make([]float64, binNum)
	hist.BinEdges = make([]float64, binNum + 1)

	hist.logHistogram = false

	hist.lowLim = low
	hist.highLim = high
	hist.binWidth = (hist.highLim - hist.lowLim) / float64(binNum)

	for i := 0; i < binNum; i++ {
		hist.BinEdges[i] = hist.lowLim + hist.binWidth * float64(i)
		hist.BinValues[i] = hist.lowLim + hist.binWidth * (float64(i) + 0.5)
	}

	hist.BinEdges[binNum] = hist.highLim

	return hist, hist.AddArray(xs)
}

// NewBoundedLogHistogram creates a Histogram instance out of the given array
// of values with the given number of bins which fall between the given limits.
// The logarithms of bin centers are uniformly dist. Any
// values outside of these limits are ignored. The returned integer is the
// number of such ignored values. Because of this, infinte and non-positive
// values do not cause a panic.
//
// NewBoundedHistogram panics if given a non-positive number of bins or
// a low bound as large or larger than the high bound or if given infinite bounds.
func NewBoundedLogHistogram(xs []float64, binNum int, low, high float64) (*Histogram, int) {
	if binNum < 1 {
		panic(fmt.Sprintf("stats.NewBoundedLogHistogram given binNum of %d", binNum))
	} else if low >= high || low <= 0 || math.IsInf(low, 0) || math.IsInf(high, 0) {
		panic(fmt.Sprintf("stats.NewBoundedLogHistogram given range [%d, %d]", low, high))
	}

	hist := new(Histogram)
	hist.Bins = make([]int, binNum)
	hist.BinValues = make([]float64, binNum)
	hist.BinEdges = make([]float64, binNum + 1)

	hist.logHistogram = true

	hist.lowLim = math.Log(low)
	hist.highLim = math.Log(high)
	hist.binWidth = (hist.highLim - hist.lowLim) / float64(binNum)

	for i := 0; i < binNum; i++ {
		hist.BinEdges[i] = math.Exp(hist.lowLim + hist.binWidth * float64(i))
		hist.BinValues[i] = math.Exp(hist.lowLim + hist.binWidth * (float64(i) + 0.5))
	}

	hist.BinEdges[binNum] = hist.highLim

	return hist, hist.AddArray(xs)
}
