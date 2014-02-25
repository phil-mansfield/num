package num

import (
	"math"
	"math/rand"
	"testing"
)

func TestFind(t *testing.T) {
	lin := func(x float64) float64 { return x }
	sqr := func(x float64) float64 { return x * x }

	tests := []struct {
		f                     Func1D
		target, min, max, exp float64
	}{
		{lin, 3.14, 0, 10, 3.14},
		{lin, 10, 0, 10, 10},
		{sqr, 9, -10, 0, -3},
	}

	for _, test := range tests {
		res := Find(test.f, test.target, test.min, test.max)
		if !almostEq(res, test.exp) {
			t.Errorf("Find(%q, %.5g, %.5g, %.5g) -> %.5g, wanted %.5g",
				test.f, test.target, test.min, test.max, res, test.exp)
		}
	}
}

func TestFindZero(t *testing.T) {
	lin := func(x float64) float64 { return x }
	sqr := func(x float64) float64 { return x*x - 9.0 }

	tests := []struct {
		f                 Func1D
		guess, scale, exp float64
	}{
		{lin, 3.14, 1.0, 0.0},
		{sqr, 4.0, 1.0, 3.0},
		{sqr, 2.0, 1.0, 3.0},
		{sqr, -4.0, 1.0, -3.0},
		{sqr, -2.0, 1.0, -3.0},
	}

	for _, test := range tests {
		res := FindZero(test.f, test.guess, test.scale)
		if !almostEq(res, test.exp) {
			t.Errorf("Find(%q, %.5g, %.5g) -> %g, wanted %.5g",
				test.f, test.guess, test.scale, res, test.exp)
		}
	}
}

func BenchmarkFindLightweight(b *testing.B) {
	sqr := func(x float64) float64 { return x*x - 50.0 }
	exps := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		exps[i] = (0.5 - rand.Float64()) * 100
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Find(sqr, exps[i], 0, 10)
	}
}

func BenchmarkFindZeroLightweight(b *testing.B) {
	sqr := func(x float64) float64 { return x*x - 50.0 }
	guesses := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		guesses[i] = rand.Float64() * 10
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FindZero(sqr, guesses[i], 1.0)
	}
}

func BenchmarkFindMediumweight(b *testing.B) {
	sqr := func(x float64) float64 {
		return math.Pow(x, 2.0) - math.Pow(math.Sqrt(50.0), 2.0)
	}
	exps := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		exps[i] = (0.5 - rand.Float64()) * 100
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Find(sqr, exps[i], 0, 10)
	}
}

func BenchmarkFindZeroMediumweight(b *testing.B) {
	sqr := func(x float64) float64 {
		return math.Pow(x, 2.0) - math.Pow(math.Sqrt(50.0), 2.0)
	}
	guesses := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		guesses[i] = rand.Float64() * 10
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FindZero(sqr, guesses[i], 1.0)
	}
}

func BenchmarkFindHeavyweight(b *testing.B) {
	sqr := func(x float64) float64 {
		sum := 0.0
		for i := 0; i < 50; i++ {
			sum += math.Pow(2.0, 2.0) / 4.0
		}
		return math.Pow(x, 2.0) - sum
	}
	exps := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		exps[i] = (0.5 - rand.Float64()) * 100
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Find(sqr, exps[i], 0, 10)
	}
}

func BenchmarkFindZeroHeavyweight(b *testing.B) {
	sqr := func(x float64) float64 {
		sum := 0.0
		for i := 0; i < 50; i++ {
			sum += math.Pow(2.0, 2.0) / 4.0
		}
		return math.Pow(x, 2.0) - sum
	}
	guesses := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		guesses[i] = rand.Float64() * 10
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FindZero(sqr, guesses[i], 1.0)
	}
}
