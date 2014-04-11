package stats

import (
	"fmt"
	"math"
)

type ArrayStats struct {
	Len int
	Mean, Variance, StdDev, Max, Min float64
}

func TrimDescribe(xs []float64, loLim, hiLim float64) *ArrayStats {
	if len(xs) <= 0 {
		panic("stats.TrimDescribe given array of length zero.")
	} else if math.IsNaN(hiLim) || math.IsNaN(loLim) || hiLim <= loLim {
		panic(fmt.Sprintf("stats.TrimDescribe given invalid bounds: [%g, %g].",
			loLim, hiLim))
	}
	
	count := 0
	sum := 0.0
	sqrSum := 0.0
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
		if xs[i] > max { max = xs[i] }
		if xs[i] < min { min = xs[i] }
	}

	if count == 0 {
		panic(fmt.Sprintf("stats.TrimDescribe given array with no elements in [%g, %g].",
			loLim, hiLim))
	}

	as := new(ArrayStats)
	as.Len = count
	as.Mean = sum / float64(as.Len)
	as.Variance = sqrSum / float64(as.Len) - as.Mean * as.Mean
	as.StdDev = math.Sqrt(as.Variance)
	as.Max = max
	as.Min = min
	return as
}

func Describe(xs []float64) *ArrayStats {
	if len(xs) <= 0 { panic("stats.Describe given array of length zero.") }
	return TrimDescribe(xs, math.Inf(-1), math.Inf(+1))
}

func Mean(xs []float64) float64 { 
	if len(xs) <= 0 { panic("stats.Mean given array of length zero.") }
	return Describe(xs).Mean
}

func Variance(xs []float64) float64 {
	if len(xs) <= 0 { panic("stats.Variance given array of length zero.") }
	return Describe(xs).Variance
}

func StdDev(xs []float64) float64 {
	if len(xs) <= 0 { panic("stats.StdDev given array of length zero.") }
	return Describe(xs).StdDev
}

func Max(xs []float64) float64 {
	if len(xs) <= 0 { panic("stats.Max given array of length zero.") }
	return Describe(xs).Max
}

func Min(xs []float64) float64 {
	if len(xs) <= 0 { panic("stats.Min given array of length zero.") }
	return Describe(xs).Min
}
