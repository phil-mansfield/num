package stats

import (
	"math"
)

type Histogram struct {
	Bins []int
	Values, BinEdges []float64
	ItemCount int

	binWidth float64
	lowLim, highLim float64
	logHistogram bool
}

func (hist *Histogram) Add(x float64) int {
	if hist.logHistogram { x = math.Log(x) }
	if x == hist.highLim { hist.Bins[len(hist.Bins) - 1]++ }
	idx := int((x - hist.lowLim) / hist.binWidth)
	if idx < 0 || idx >= len(hist.Bins) { return 1 }
	hist.ItemCount++
	hist.Bins[idx]++
	return 0
}

func (hist *Histogram) AddArray(xs []float64) int {
	initialItemCount := 0

	for _, x := range xs {
		if hist.logHistogram { x = math.Log(x) }
		if x == hist.highLim { hist.Bins[len(hist.Bins) - 1]++ }
		idx := int((x - hist.lowLim) / hist.binWidth)
		if idx < 0 || idx >= len(hist.Bins) { continue }
		hist.ItemCount++
		hist.Bins[idx]++
	}
	
	return hist.ItemCount - initialItemCount + len(xs)
}

func (hist *Histogram) NormalizedBins(area float64) []float64 {
	itemArea := area / float64(hist.ItemCount)
	
	normedBins := make([]float64, len(hist.Bins))
	for i := 0; i < len(hist.Bins); i++ {
		normedBins[i] = float64(hist.Bins[i]) * itemArea
	}

	return normedBins
}

func (hist *Histogram) CumulativeBins() []int {
	cumBins := make([]int, len(hist.Bins))
	binNum := len(hist.Bins)
	
	cumBins[binNum - 1] = hist.Bins[binNum - 1]
	for i := binNum - 2; i >= 0; i-- {
		cumBins[i] = cumBins[i + 1] + hist.Bins[i]
	}

	return cumBins
}

func NewHistogram(xs []float64, binNum int) *Histogram {
	as := Describe(xs)
	min, max := as.Min, as.Max
	hist, _ := NewBoundedHistogram(xs, binNum, min, max)
	return hist
}

func NewLogHistogram(xs []float64, binNum int) *Histogram {
	as := Describe(xs)
	min, max := as.Min, as.Max
	hist, _ := NewBoundedLogHistogram(xs, binNum, min, max)
	return hist
}

func NewBoundedHistogram(xs []float64, binNum int, low, high float64) (*Histogram, int) {
	hist := new(Histogram)
	hist.Bins = make([]int, binNum)
	hist.Values = make([]float64, binNum)
	hist.BinEdges = make([]float64, binNum + 1)

	hist.logHistogram = false

	hist.lowLim = low
	hist.highLim = high
	hist.binWidth = (hist.highLim - hist.lowLim) / float64(binNum)

	for i := 0; i < binNum; i++ {
		hist.BinEdges[i] = hist.lowLim + hist.binWidth * float64(i)
		hist.Values[i] = hist.lowLim + hist.binWidth * (float64(i) + 0.5)
	}

	hist.BinEdges[binNum] = hist.highLim

	return hist, hist.AddArray(xs)
}

func NewBoundedLogHistogram(xs []float64, binNum int, low, high float64) (*Histogram, int) {
	hist := new(Histogram)
	hist.Bins = make([]int, binNum)
	hist.Values = make([]float64, binNum)
	hist.BinEdges = make([]float64, binNum + 1)

	hist.logHistogram = true

	hist.lowLim = math.Log(low)
	hist.highLim = math.Log(high)
	hist.binWidth = (hist.highLim - hist.lowLim) / float64(binNum)

	for i := 0; i < binNum; i++ {
		hist.BinEdges[i] = hist.lowLim + hist.binWidth * float64(i)
		hist.Values[i] = hist.lowLim + hist.binWidth * (float64(i) + 0.5)
	}

	hist.BinEdges[binNum] = hist.highLim

	return hist, hist.AddArray(xs)
}
