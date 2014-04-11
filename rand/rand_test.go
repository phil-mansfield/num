package rand

import (
	"testing"
	
	"github.com/phil-mansfield/num/objects/geom"
)

// This really sucks. Rewrite all of it.

func BenchmarkTauswortheNext(b *testing.B) {
	rand := NewTimeSeed(Tausworthe)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkGoRandNext(b *testing.B) {
	rand := NewTimeSeed(Golang)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkXorshiftNext(b *testing.B) {
	rand := NewTimeSeed(Xorshift)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkMultiplyWithCarryNext(b *testing.B) {
	rand := NewTimeSeed(MultiplyWithCarry)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkGslNext(b *testing.B) {
	rand := NewTimeSeed(Gsl)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkSysNext(b *testing.B) {
	rand := NewTimeSeed(Sys)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkTauswortheGaussian(b *testing.B) {
	rand := NewTimeSeed(Tausworthe)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Gaussian(2, 2)
	}
}

func BenchmarkSphereAt(b *testing.B) {
	rand := NewTimeSeed(Tausworthe)

	b.ReportAllocs()
	b.ResetTimer()

	target := make([]float64, 3)
	for i := 0; i < b.N; i++ {
		rand.SphereAt(2, target)
	}
}

func BenchmarkFiniteLineAt(b *testing.B) {
	rand := NewTimeSeed(Xorshift)

	b.ReportAllocs()
	b.ResetTimer()

	target := make([]float64, 3)
	anchor := []float64{1, 1, 1}
	norm := []float64{1, 2, 1}
	line := geom.NewFiniteLine(anchor, norm, 3)
	for i := 0; i < b.N; i++ {
		rand.FiniteLineAt(line, target)
	}
}

func randomEnough(chis []float64) bool {
	badCount := 0
	for _, chi := range chis {
		if chi < 0.05 || chi > 0.95 {
			badCount += 1
		}
	}
	return badCount < 3
}

func TestFrequncyTausworthe(t *testing.T) {
	rand := NewTimeSeed(Tausworthe)
	chis, ps := FrequencyTest(rand, 12, 3 * 7 * 1000 * 1000, 7)
	if !randomEnough(ps) {
		t.Errorf("Got p(Chi^2) values of %v\n Chi^2 of", ps, chis)
	}
}

func TestFrequncyGoRand(t *testing.T) {
	rand := NewTimeSeed(Golang)
	chis, ps := FrequencyTest(rand, 12, 7 * 3 * 1000 * 1000, 7)
	if !randomEnough(ps) {
		t.Errorf("Got p(Chi^2) values of %v\n Chi^2 of", ps, chis)
	}
}

func TestFrequncyXorshift(t *testing.T) {
	rand := NewTimeSeed(Xorshift)
	chis, ps := FrequencyTest(rand, 12, 7 * 3 * 1000 * 1000, 7)
	if !randomEnough(ps) {
		t.Errorf("Got p(Chi^2) values of %v\n Chi^2 of", ps, chis)
	}
}

func TestFrequncyMultiplyWithCarry(t *testing.T) {
	rand := NewTimeSeed(MultiplyWithCarry)
	chis, ps := FrequencyTest(rand, 12, 7 * 3 * 1000 * 1000, 7)
	if !randomEnough(ps) {
		t.Errorf("Got p(Chi^2) values of %v\n Chi^2 of", ps, chis)
	}
}

func TestSerialTausworthe(t *testing.T) {
	rand := NewTimeSeed(Tausworthe)
	ps, chis := SerialTest(rand, 5, 25 * 7 * 1000 * 1000, 7)
	if !randomEnough(chis) {
		t.Errorf("Got p(Chi^2) values of %v\n Chi^2 of", ps, chis)
	}
}

func TestSerialGoRand(t *testing.T) {
	rand := NewTimeSeed(Golang)
	ps, chis := SerialTest(rand, 5, 25 * 7 * 1000 * 1000, 7)
	if !randomEnough(chis) {
		t.Errorf("Got p(Chi^2) values of %v\n Chi^2 of", ps, chis)
	}
}

func TestSerialXorshift(t *testing.T) {
	rand := NewTimeSeed(Xorshift)
	ps, chis := SerialTest(rand, 5, 25 * 7 * 1000 * 1000, 7)
	if !randomEnough(chis) {
		t.Errorf("Got p(Chi^2) values of %v\n Chi^2 of", ps, chis)
	}
}

func TestSerialMultiplyWithCarry(t *testing.T) {
	rand := NewTimeSeed(MultiplyWithCarry)
	ps, chis := SerialTest(rand, 5, 25 * 7 * 1000 * 1000, 7)
	if !randomEnough(chis) {
		t.Errorf("Got p(Chi^2) values of %v\n Chi^2 of", ps, chis)
	}
}
