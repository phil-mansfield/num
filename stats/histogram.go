package stats

import (
	"fmt"
	"math"
)

// TODO: write RemoveArray

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
	logHistogram, init bool
}

// Index returns the index that a given value would be assigned to if it
// were added to the histogram. If the given value is outside the range of the
// Histogram, -1 is returned.
func (hist *Histogram) Index(x float64) int {
	if !hist.init {
		panic("stats.Histogram.Index called on unitinitalized struct.")
	}

	var idx int

	if math.IsNaN(x) { panic("stats.Histogram.Index given NaN.") }
	if hist.logHistogram {
		if x <= 0 { return -1 }
		x = math.Log(x)
	}

	if x == hist.highLim {
		idx = len(hist.Bins) - 1
	} else {
		idx = int((x - hist.lowLim) / hist.binWidth)
		if idx < 0 || idx >= len(hist.Bins) { return -1 }
	}

	return idx
}

// Add adds a single value to the Histogram. If the value is out of bounds it
// is ignored and 1 is returned. Otherwise 0 is returned.
func (hist *Histogram) Add(x float64) int {
	if !hist.init {
		panic("stats.Histogram.Add called on unitinitalized struct.")
	}

	var idx int

	if math.IsNaN(x) { panic("stats.Histogram.Add given NaN.") }
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

// Remove removes a single element from  the Histogram. If the value is out of
// bounds or if the bin it would be removed from would become negative, the value
// is ignored and 1 is returned. Otherwise 0 is returned.
func (hist *Histogram) Remove(x float64) int {
	if !hist.init {
		panic("stats.Histogram.Remove called on unitinitalized struct.")
	}

	var idx int

	if math.IsNaN(x) { panic("stats.Histogram.Remove given NaN.") }
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

	if hist.Bins[idx] == 0 { return 1 }

	hist.ValueCount--
	hist.Bins[idx]--
	return 0
}

// AddArray adds every element of a given array to the Histogram. Elements out
// of bounds are ignored. The return value is the number of such ignored
// elements.
func (hist *Histogram) AddArray(xs []float64) int {
	if !hist.init {
		panic("stats.Histogram.AddArray called on unitinitalized struct.")
	}

	var idx int
	initialValueCount := 0

	for i, x := range xs {
		_ = i
		if math.IsNaN(x) { 
		 	panic(fmt.Sprintf("stats.Histogram.AddArray given NaN at index %d", i))
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
	
	return initialValueCount + len(xs) - hist.ValueCount
}

// InBounds returns true if the given value falls between the inclusive bounds
// of the Histogram.
func (hist *Histogram) InBounds(x float64) bool {
	if !hist.init {
		panic("stats.Histogram.InBounds called on uninitialized struct.")
	} else if math.IsNaN(x) {
		panic("stats.Histogram.InBounds given NaN.")
	}

	// This is kinda slow. Maybe we should save the non-log limits for both
	// types.
	if hist.logHistogram {
		x = math.Log(x)
	}

	return x <= hist.highLim && x >= hist.lowLim
}

// NormalizedBins returns an array with the same relative frequencies as
// hist.Bins, but whose total sum is equal to a given area.
func (hist *Histogram) NormalizedBins(area float64) []float64 {
	if !hist.init {
		panic("stats.Histogram.NormalizedBins called on unitinitalized struct.")
	}

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
	if !hist.init {
		panic("stats.Histogram.CumulativeBins called on unitinitalized struct.")
	}

	cumBins := make([]int, len(hist.Bins))
	binNum := len(hist.Bins)
	
	cumBins[binNum - 1] = hist.Bins[binNum - 1]
	for i := binNum - 2; i >= 0; i-- {
		cumBins[i] = cumBins[i + 1] + hist.Bins[i]
	}

	return cumBins
}

// InitHistogram initializes a Histogram instance from the given array of values
// with the given number of bins. The minimum and maximum allowed values are
// taken to be the minimum and maximum values in that array.
//
// The returned value is the initialized Histogram.
//
// Init will panic if given a non-positive number of bins or less than
// two starting values, as either option leads to an ill-defined range.
// Init will panic fi given any infinite values for the same reason.
func (hist *Histogram) Init(xs []float64, binNum int) *Histogram {
	if hist.init {
		panic("stats.Histogram.Init called on initialized struct.")
	} else if len(xs) <= 1 {
		panic("stats.Histogram.Init given empty array.")
	} else if binNum < 1 {
		panic(fmt.Sprintf("stats.Histogram.Init given binNum of %d", binNum))
	}

	as := Describe(xs)
	min, max := as.Min, as.Max

	if math.IsInf(min, 0) {
		panic(fmt.Sprintf("stats.Histogram.Init given array with infinite value at %d", as.MinIdx))
	} else if math.IsInf(max, 0)  {
		panic(fmt.Sprintf("stats.Histogram.Init given array with infinite value at %d", as.MaxIdx))
	}

	hist.InitBounded(xs, binNum, min, max)
	return hist
}

// InitLog initializes a Histogram instance from the given array of values
// with the given number of bins. Logarithms of bin centers are uniformly
// distributed. The minimum and maximum allowed values are taken to be the
// minimum and maximum values in that array.
//
// The returned value is the initialized Histogram.
//
// InitLog will panic if given a non-positive number of bins or less than
// two starting values, as either option leads to an ill-defined range. InitLog
// will also panic if given any negative values or infinite values for the same reason.
func (hist *Histogram) InitLog(xs []float64, binNum int) *Histogram {
	if hist.init {
		panic("stats.Histogram.InitLog called on initialized struct.")
	} else if len(xs) <= 1 {
		panic("stats.Histogram.InitLog given empty array.")
	} else if binNum < 1 {
		panic(fmt.Sprintf("stats.Histogram.InitLog given binNum of %d", binNum))
	}

	as := Describe(xs)
	min, max := as.Min, as.Max
	
	if min <= 0 {
		panic(fmt.Sprintf("stats.Histogram.InitLog given non-positive value %d at %d", min, as.MinIdx))
	}

	hist.InitBoundedLog(xs, binNum, min, max)
	return hist
}

// InitBounded initializes a Histogram instance from the given array
// of values with the given number of bins which fall between the given
// limits. Any values outside of these limits are ignored. The returned
// integer is the number of such ignored values. Because of this infinite values
// do not cause a panic.
//
// The first returned value is the initialized Histogram.
//
// InitBounded panics if given a non-positive number of bins or
// a low bound as large or larger than the high bound or if given infinte
// bounds.
func (hist *Histogram) InitBounded(xs []float64, binNum int, low, high float64) (*Histogram, int) {
	if hist.init {
		panic("stats.Histogram.InitBounded called on initialized struct.")
	} else 	if binNum < 1 {
		panic(fmt.Sprintf("stats.Histogram.InitBounded given binNum of %d",
			binNum))
	} else if low >= high || math.IsInf(low, 0) || math.IsInf(high, 0) ||
		math.IsNaN(low) || math.IsNaN(high) {
		panic(fmt.Sprintf("stats.Histogram.InitBounded given range [%d, %d]",
			low, high))
	}

	hist.init = true
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

// InitBoundedLog initializes a Histogram instance from the given array
// of values with the given number of bins which fall between the given limits.
// The logarithms of bin centers are uniformly dist. Any
// values outside of these limits are ignored. The returned integer is the
// number of such ignored values. Because of this, infinte and non-positive
// values do not cause a panic.
//
// The first returned value is the initialized Histogram.
//
// InitBoundedLog panics if given a non-positive number of bins or
// a low bound as large or larger than the high bound or if given infinite bounds.
func (hist *Histogram) InitBoundedLog(xs []float64, binNum int, low, high float64) (*Histogram, int) {
	if hist.init {
		panic("stats.Histogram.InitBoundedLog called on initialized struct.")
	} else if binNum < 1 {
		panic(fmt.Sprintf("stats.Histogram.InitBoundedLog given binNum of %d", binNum))
	} else if low >= high || low <= 0 || math.IsInf(low, 0) ||
		math.IsInf(high, 0) || math.IsNaN(low) || math.IsNaN(high) {
		panic(fmt.Sprintf("stats.Histogram.InitBoundedLog given range [%d, %d]", low, high))
	}

	hist.init = true
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
