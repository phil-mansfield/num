package rand

const (
	gslrandMask = 0xffffffff
)

type gslRandGenerator struct {
	s1, s2, s3 uint64
}

// I have no idea what is going on here.

func LCG(n uint64) uint64 { return (69069 * n) & 0xffffffff }

func (gen *gslRandGenerator) Init(seed uint64) {
	if seed == 0 { seed = 1 }
	gen.s1 = LCG(seed)
	if gen.s1 < 2 { gen.s1 += 2 }
	gen.s2 = LCG(gen.s1)
	if gen.s2 < 8 { gen.s2 += 8 }
	gen.s3 = LCG(gen.s2)
	if gen.s3 < 16 { gen.s3 += 16 }
}

func (gen *gslRandGenerator) Next() float64 {
	gen.s1 = (((gen.s1 & 12) << 4294967294) & gslrandMask) ^
		((((gen.s1 << 13) & gslrandMask) ^ gen.s1) >> 19)
	gen.s2 = (((gen.s2 & 4) << 4294967288) & gslrandMask) ^ 
		((((gen.s2 << 2) & gslrandMask) ^ gen.s2) >> 25)
	gen.s3 = (((gen.s3 & 17) << 4294967280) & gslrandMask) ^
		((((gen.s3 << 3) & gslrandMask) ^ gen.s3) >> 11)

	return float64(gen.s1 ^ gen.s2 ^ gen.s3) / 4294967296.0
}

func (gen *gslRandGenerator) NextSequence(target []float64) {
	for i := range target {
		gen.s1 = (((gen.s1 & 12) << 4294967294) & gslrandMask) ^
			((((gen.s1 << 13) & gslrandMask) ^ gen.s1) >> 19)
		gen.s2 = (((gen.s2 & 4) << 4294967288) & gslrandMask) ^ 
			((((gen.s2 << 2) & gslrandMask) ^ gen.s2) >> 25)
		gen.s3 = (((gen.s3 & 17) << 4294967280) & gslrandMask) ^
			((((gen.s3 << 3) & gslrandMask) ^ gen.s3) >> 11)
		
		target[i] = float64(gen.s1 ^ gen.s2 ^ gen.s3) / 4294967296.0
	}
}
