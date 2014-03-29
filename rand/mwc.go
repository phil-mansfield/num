package rand

import (
	"math"
)

const (
	mwcPhi uint32 = 0x9e3779b9
	mwcA uint64 = 18782
	mwcR uint32 = 0xfffffffe
)

var (
	mwcMaxUint = float64(math.MaxUint32)
)

type multiplyWithCarryGenerator struct {
	seq [4096]uint32
	c uint32
	i int32
}

func (gen *multiplyWithCarryGenerator) Init(seed uint64) {
	x := uint32(seed)
	gen.seq[0] = x
	gen.seq[1] = x + mwcPhi
	gen.seq[2] = x + mwcPhi + mwcPhi

	for gen.i = 3; gen.i < 4096; gen.i++ {
		gen.seq[gen.i] = gen.seq[gen.i - 3] ^ 
			gen.seq[gen.i - 2] ^ mwcPhi ^ uint32(gen.i)
	}

	gen.c = 362463
}

func (gen *multiplyWithCarryGenerator) Next() float64 {
	gen.i = (gen.i + 1) & 4095

	t := mwcA * uint64(gen.seq[gen.i]) + uint64(gen.c)
	gen.c = uint32(t >> 32)
	x := uint32(t + uint64(gen.c))
	if x < gen.c {
		x++
		gen.c++
	}
	gen.seq[gen.i] = mwcR - x
	return float64(gen.seq[gen.i]) / mwcMaxUint
}
