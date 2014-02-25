package gorecipies

// ChiSqrDist computes cumulative probability that Chi^2 could
// take on a value less than chiSqr with the given number of degrees
// of freedom.
func ChiSqrDist(chiSqr float64, nu int) float64 {
	if nu < 1 {
		panic("Less than one degree of freedom in ChiSqrDist.")
	}

	return IncGamma(float64(nu) / 2, chiSqr / 2)
}

// ChiSqr calculates the Chi^2 value and the cumulative pdf of that value
// occuring for a set of outcomes which have a probability probs[i] of
// occuring and have occured counts[i] times.
func ChiSqr(counts []int64, probs []float64) (chiSqr, prob float64) {
	if len(counts) <= 1 {
		panic("|counts| <= 1 in ChiSqr.")
	} else if len(counts) != len(probs) {
		panic("|counts| != |probs| in ChiSqr")
	}

	n := int64(0)
	for _, count := range counts {
		n += count
	}

	sqrSum := 0.0

	for i, count := range counts {
		sqrSum += float64(count * count) / probs[i]
	}

	chiSqr = sqrSum / float64(n) - float64(n)
	prob = ChiSqrDist(chiSqr, len(counts) - 1)

	return chiSqr, prob
}
