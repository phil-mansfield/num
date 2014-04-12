package rand

import (
	"testing"
)

func benchmarkUniform(gt GeneratorType, b *testing.B) {
	gen := NewTimeSeed(gt)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = gen.Uniform(0, 13)
	}
}

func benchmarkUniformAt(gt GeneratorType, tLen int, b *testing.B) {
	gen := NewTimeSeed(gt)
	b.ResetTimer()

	target := make([]float64, tLen)

	n := 0
	for n < b.N {
		if n + tLen > b.N { target = target[0: b.N - n] }
		gen.UniformAt(0, 13, target)
		n += tLen
	}
}


func BenchmarkUniformGsl(b *testing.B) { benchmarkUniform(Gsl, b) }
func BenchmarkUniformGolang(b *testing.B) { benchmarkUniform(Golang, b) }
func BenchmarkUniformMwc(b *testing.B) { benchmarkUniform(MultiplyWithCarry, b) }
func BenchmarkUniformXorshift(b *testing.B) { benchmarkUniform(Xorshift, b) }
func BenchmarkUniformTausworthe(b *testing.B) { benchmarkUniform(Tausworthe, b) }

func BenchmarkUniformAtGsl(b *testing.B) { benchmarkUniformAt(Gsl, DefaultBufSize, b) }
func BenchmarkUniformAtGolang(b *testing.B) { benchmarkUniformAt(Golang, DefaultBufSize, b) }
func BenchmarkUniformAtMwc(b *testing.B) { benchmarkUniformAt(MultiplyWithCarry, DefaultBufSize, b) }
func BenchmarkUniformAtXorshift(b *testing.B) { benchmarkUniformAt(Xorshift, DefaultBufSize, b) }
func BenchmarkUniformAtTausworthe(b *testing.B) { benchmarkUniformAt(Tausworthe, DefaultBufSize, b) }
