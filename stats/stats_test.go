package stats

import (
	"sort"
	"testing"

	"github.com/phil-mansfield/num"
	"github.com/phil-mansfield/num/rand"
)

func floatArrayEq(xs, ys []float64) bool {
	if len(xs) != len(ys) { return false }
	for i := 0; i < len(xs); i++ {
		if !num.AlmostEqual(xs[i], ys[i]) { return false }
	}
	return true
}

func intArrayEq(xs, ys []int) bool {
	if len(xs) != len(ys) { return false }
	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] { return false }
	}
	return true
}

func consistentValueCount(hist *Histogram) bool {
	sum := 0
	for i := 0; i < len(hist.Bins); i++ {
		sum += hist.Bins[i]
	}
	return sum == hist.ValueCount
}

func TestHistogramNewBounded(t *testing.T) {
	empty := make([]float64, 0)

	targetValues := []float64{-4.5, -3.5, -2.5, -1.5, -0.5,
		+0.5, +1.5, +2.5, +3.5, +4.5}
	targetEdges := []float64{-5, -4, -3, -2, -1,
		+0, +1, +2, +3, +4, +5}
	targetBins := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	targetMisses := 0

	hist, misses := new(Histogram).InitBounded(empty, 10, -5, +5)

	if !floatArrayEq(targetValues, hist.BinValues) { 
		t.Errorf("Histogram has wrong values: %v, want %v",
			hist.BinValues, targetValues)
	} else if !floatArrayEq(targetEdges, hist.BinEdges) {
		t.Errorf("Histogram has wrong edges: %v, want %v",
			hist.BinEdges, targetEdges)
	} else if !intArrayEq(targetBins, hist.Bins) {
		t.Errorf("Histogram has wrong bins: %v, want %v",
			hist.Bins, targetBins)
	} else if !consistentValueCount(hist) {
		t.Errorf("Histogram has inconsistent value count: %v, has %v",
			hist.ValueCount, hist.Bins)
	} else if misses != targetMisses {
		t.Errorf("Histogram generated %d misses, want %d",
			misses, targetMisses)
	}
}

func TestHistogramAddHit(t *testing.T) {
	empty := make([]float64, 0)

	targetBins := []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 0}
	targetMisses := 0

	hist, misses := new(Histogram).InitBounded(empty, 10, -5, +5)

	hist.Add(-3)

	if !intArrayEq(targetBins, hist.Bins) {
		t.Errorf("Histogram has wrong bins: %v, want %v",
			hist.Bins, targetBins)
	} else if !consistentValueCount(hist) {
		t.Errorf("Histogram has inconsistent value count: %v, has %v",
			hist.ValueCount, hist.Bins)
	} else if misses != targetMisses {
		t.Errorf("Histogram generated %d misses, want %d",
			misses, targetMisses)
	}
}

func TestHistogramAddMaximum(t *testing.T) {
	empty := make([]float64, 0)

	targetBins := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	targetMisses := 0

	hist, misses := new(Histogram).InitBounded(empty, 10, -5, +5)

	misses += hist.Add(5)

	if !intArrayEq(targetBins, hist.Bins) {
		t.Errorf("Histogram has wrong bins: %v, want %v",
			hist.Bins, targetBins)
	} else if !consistentValueCount(hist) {
		t.Errorf("Histogram has inconsistent value count: %v, has %v",
			hist.ValueCount, hist.Bins)
	} else if misses != targetMisses {
		t.Errorf("Histogram generated %d misses, want %d",
			misses, targetMisses)
	}
}


func TestHistogramAddMinimum(t *testing.T) {
	empty := make([]float64, 0)

	targetBins := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	targetMisses := 0
	
	hist, misses := new(Histogram).InitBounded(empty, 10, -5, +5)

	misses += hist.Add(-5)
	
	if !intArrayEq(targetBins, hist.Bins) {
		t.Errorf("Histogram has wrong bins: %v, want %v",
			hist.Bins, targetBins)
	} else if !consistentValueCount(hist) {
		t.Errorf("Histogram has inconsistent value count: %v, has %v",
			hist.ValueCount, hist.Bins)
	} else if misses != targetMisses {
		t.Errorf("Histogram generated %d misses, want %d",
			misses, targetMisses)
	}
}

func TestHistogramAddMiss(t *testing.T) {
	empty := make([]float64, 0)
	
	targetBins := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	targetMisses := 2

	hist, misses := new(Histogram).InitBounded(empty, 10, -5, +5)

	misses += hist.Add(6)
	misses += hist.Add(-6)

	if !intArrayEq(targetBins, hist.Bins) {
		t.Errorf("Histogram has wrong bins: %v, want %v",
			hist.Bins, targetBins)
	} else if !consistentValueCount(hist) {
		t.Errorf("Histogram has inconsistent value count: %v, has %v",
			hist.ValueCount, hist.Bins)
	} else if misses != targetMisses {
		t.Errorf("Histogram generated %d misses, want %d",
			misses, targetMisses)
	}
}

func TestHistogramAddArray(t *testing.T) {
	empty := make([]float64, 0)
	
	targetBins := []int{1, 1, 0, 0, 0, 0, 0, 0, 0, 1}
	targetMisses := 2
	
	values := []float64{-5, 5, -3.5, -6, 6}

	hist, misses := new(Histogram).InitBounded(empty, 10, -5, +5)

	misses += hist.AddArray(values)

	if !intArrayEq(targetBins, hist.Bins) {
		t.Errorf("Histogram has wrong bins: %v, want %v",
			hist.Bins, targetBins)
	} else if !consistentValueCount(hist) {
		t.Errorf("Histogram has inconsistent value count: %v, has %v",
			hist.ValueCount, hist.Bins)
	} else if misses != targetMisses {
		t.Errorf("Histogram generated %d misses, want %d",
			misses, targetMisses)
	}
}

func BenchmarkAddArrayTwice(b *testing.B) {
	input := make([]float64, b.N)
	gen := rand.NewTimeSeed(rand.Xorshift)
	gen.UniformAt(0, 100, input)

	hist, _ := new(Histogram).InitBounded([]float64{}, 1024, 0, 100)

	b.ResetTimer()
	hist.AddArray(input)
	hist.AddArray(input)
}

func BenchmarkAddTwice(b *testing.B) {
	input := make([]float64, b.N)
	gen := rand.NewTimeSeed(rand.Xorshift)
	gen.UniformAt(0, 100, input)

	hist, _ := new(Histogram).InitBounded([]float64{}, 1024, 0, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hist.Add(input[i])
	}
	for i := 0; i < b.N; i++ {
		hist.Add(input[i])
	}
}

func BenchmarkSort(b *testing.B) {
	input := make([]float64, b.N)
	gen := rand.NewTimeSeed(rand.Xorshift)
	gen.UniformAt(0, 100, input)

	b.ResetTimer()

	sort.Float64s(input)
}
