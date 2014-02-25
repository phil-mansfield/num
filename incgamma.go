package gorecipies

import (
	"fmt"
	"math"
)

// incGammaContinuedFraction computes the incomplete gamma function
// via the contiued fraction:
//
// exp(-x) x^ a (1 / (x + 1 - a - (1 (1 - a)) / (x + 3 - a - ...))).
//
// This is implemented via modified Lentz's method.
func incGammaContinuedFraction(a, x float64) float64 {
	minValue := math.SmallestNonzeroFloat64

	// Start computing terms at n = 1.

	// c_n = A_n / A_n-1. Note that A_0 = a_0 = 0.
	c := math.MaxFloat64

	// c_n = B_n-1 / B_n. Note that B_0 = b_0 = 1.
	b := x + 1 - a
	d := 1 / b
	
	fracEst := d

	for n := 1; n <= ConvergenceIters; n++ {
		an := -float64(n) * (float64(n) - a)
		b += 2.0

		d = an * d + b
		if math.Abs(d) < minValue { d = minValue }

		c = an / c + b
		if math.Abs(c) < minValue { c = minValue }

		d = 1 / d
		diff := d*c
		fracEst *= diff

		if CloseEnough(1, diff - 1) {
			lg , _ := math.Lgamma(a)
			return fracEst * math.Exp(-x + a * math.Log(x) - lg)
		}
	}

	fmtStr := "incGammaContinuedFraction(%g, %g) failed to converge"
	panic(fmt.Sprintf(fmtStr, a, x))
}

// incGammaSeries computes the incomplete gamma function via the
// series sum:
//
// gamma(a, x) = exp(-x) x^a sum_n (Gamma(a) / Gamma(a + 1 + n)) x^n.
func incGammaSeries(a, x float64) float64 {
	an := a
	termVal := 1 / an
	sum := 1 / an

	for n := 0; n < ConvergenceIters; n++ {
		an++
		termVal *= x / float64(an)
		sum += termVal
		if CloseEnough(sum, termVal) {
			// Note that x^a = exp(a * log(x)).
			lg, _ := math.Lgamma(a)
			return sum * math.Exp(-x + a * math.Log(x) - lg)
		}
	}

	fmtStr := "incGammaSeries(%g, %g) failed to converge"
	panic(fmt.Sprintf(fmtStr, a, x))
}

// IncompleteGamma computes the incomplete gamma function,
//
// P(a, x) =  gamma(a, x)/ Gamma(a),
//
// gamma(a, x) = int_0^x dt t^(a-1) exp(-t),
// Gamma(a) = int_0^inf dt t^(a-1) exp(-t).
func IncGamma(a, x float64) float64 {
	if x < 0 || a <= 0 {
		fmtStr := "x = %g, a = %g invlaid inputs for PGamma(a, x)"
		panic(fmt.Sprintf(fmtStr, x, a))
	}

	if x < a + 1 {
		return 1 - incGammaSeries(a, x)
	} else { // x >= a + 1
		// Continued fraction converges more quickly in this range
		return 1 - incGammaContinuedFraction(a, x)
	}
}
