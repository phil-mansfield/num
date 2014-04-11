package rand

import (
	"math"
)

var (
	xorshiftMaxUint = float64(math.MaxUint32)
)

type xorshiftGenerator struct {
	w, x, y, z uint32
}

func (gen *xorshiftGenerator) Init(seed uint64) {
	gen.x = 123456789
	gen.y = 362436069
	gen.z = 521288629
	gen.w = uint32(seed)
}

func (gen *xorshiftGenerator) Next() float64 {
	t := gen.x ^ (gen.x << 11)
	gen.x, gen.y, gen.z = gen.y, gen.z, gen.w
	gen.w = gen.w ^ (gen.w >> 19) ^ (t ^ (t >> 8))
	res := float64(math.MaxUint32 - gen.w) / xorshiftMaxUint
	if res == 1.0 { return gen.Next() }
	return res
}
