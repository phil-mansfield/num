package rand

import (
	"github.com/phil-mansfield/num/stats"
)

func frequencyTest(gen *Generator, d, iters int) (float64, float64) {
	bins := make([]int, d)
	for n := 0; n < iters; n++ {
		bins[gen.UniformInt(0, d - 1)]++
	}

	probs := make([]float64, d)
	for i, _ := range probs {
		probs[i] = 1 / float64(d)
	}

	chiSqr, prob := stats.ChiSqr(bins, probs)
	return chiSqr, prob
}

func FrequencyTest(gen *Generator, d, iters, chiNum int) ([]float64, []float64) {
	ps := make([]float64, chiNum)
	chis := make([]float64, chiNum)

	for i := 0; i < chiNum; i++ {
		chis[i], ps[i] = frequencyTest(gen, d, iters / chiNum)
	}

	return chis, ps
}

func serialTest(gen *Generator, d, iters int) (float64, float64) {
	bins := make([]int, d * d)
	for n := 0; n < iters; n++ {
		x, y := gen.UniformInt(0, d - 1), gen.UniformInt(0, d - 1)
		bins[y * d + x]++
	}

	probs := make([]float64, d * d)
	for i, _ := range probs {
		probs[i] = 1 / float64(d * d)
	}

	chi, prob := stats.ChiSqr(bins, probs)
	return chi, prob
}

func SerialTest(gen *Generator, d, iters, chiNum int) ([]float64, []float64) {
	chis := make([]float64, chiNum)
	probs := make([]float64, chiNum)

	for i := 0; i < chiNum; i++ {
		chis[i], probs[i] = serialTest(gen, d, iters / chiNum)
	}

	return chis, probs
}
