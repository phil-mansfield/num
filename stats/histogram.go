package stats

import (
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
	// BinEdges[i] is the lower limit of Bin[i] and the upper limit of
	// Bin[i - 1].
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

	if hist.logHistogram { x = math.Log(x) }

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

	for _, x := range xs {
		if hist.logHistogram { x = math.Log(x) }

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
func NewHistogram(xs []float64, binNum int) *Histogram {
	as := Describe(xs)
	min, max := as.Min, as.Max
	hist, _ := NewBoundedHistogram(xs, binNum, min, max)
	return hist
}

// NewHistogram creates a Histogram instance out of the given array of values
// with the given number of bins. Logarithms of bin centers are uniformly
// distributed. The minimum and maximum allowed values are taken to be the
// minimum and maximum values in that array.
func NewLogHistogram(xs []float64, binNum int) *Histogram {
	as := Describe(xs)
	min, max := as.Min, as.Max
	hist, _ := NewBoundedLogHistogram(xs, binNum, min, max)
	return hist
}

// NewHistogram creates a Histogram instance out of the given array of values
// with the given number of bins which fall between the given limits. Any
// values outside of these limits are ignored. The returned integer is the
// number of such ignored values.
func NewBoundedHistogram(xs []float64, binNum int, low, high float64) (*Histogram, int) {
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

// NewHistogram creates a Histogram instance out of the given array of values
// with the given number of bins which fall between the given limits. The
// logarithms of bin centers are uniformly dist. Any
// values outside of these limits are ignored. The returned integer is the
// number of such ignored values.
func NewBoundedLogHistogram(xs []float64, binNum int, low, high float64) (*Histogram, int) {
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
