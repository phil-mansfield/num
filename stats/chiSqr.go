package stats

import (
	"fmt"

	"github.com/phil-mansfield/num"
)

// ChiSqrDist computes cumulative probability that Chi^2 could
// take on a value less than chiSqr with the given number of degrees
// of freedom, nu.
//
// ChiSqrDist panics if given non-positive Chi^2 or nu values.
func ChiSqrDist(chiSqr float64, nu int) float64 {
	if nu <= 0 {
		panic(fmt.Sprintf("stats.ChiSqrDist given %d degrees of freedom.", nu))
	} else if chiSqr <= 0 {
		panic(fmt.Sprintf("stats.ChiSqrDist given %g degrees of freedom.", chiSqr))
	}

	return num.IncGamma(float64(nu) / 2, chiSqr / 2)
}

// ChiSqr calculates the Chi^2 value and the probability of a value less
// than this occuring for a set of outcomes which have a probability probs[i] of
// occuring and have occured counts[i] times.
//
// Most of the time, if prob is within 0.05 of 0 or 1 a distribution can be considered
// "suspiscious." If prob is within 0.01 of 0 ot 1 a distribution can be considered
// "very suspiscious." But don't tell the statisticians that I said this.
//
// ChiSqr panics if given arrays containing NaNs, if given arrays of different
// lengths, or if given arrays without at least two elements in them.
func ChiSqr(counts []int, probs []float64) (chiSqr, prob float64) {
	if len(counts) != len(probs) {
		panic(fmt.Sprintf("stats.ChiSqr given counts of length %d and probs of length %d.",
			len(counts), len(probs)))
	} else if len(counts) < 2 {
		panic(fmt.Sprintf("stats.ChiSqr given arrays of length %d", len(counts)))
	}

	n := 0
	for _, count := range counts { n += count }

	sqrSum := 0.0
	for i, count := range counts {
		sqrSum += float64(count * count) / probs[i]
	}

	chiSqr = sqrSum / float64(n) - float64(n)
	prob = ChiSqrDist(chiSqr, len(counts) - 1)

	return chiSqr, prob
}
