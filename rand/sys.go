package rand

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
)

const (
	sysBufLen = 1 << 10
)

var (
	sysMaxUint64 = float64(math.MaxUint64)
)

type sysRandGenerator struct {
	buf []byte
	idx int
}

func (gen *sysRandGenerator) Init(seed uint64) {
	gen.buf = make([]byte, 8 * sysBufLen)
	gen.idx = sysBufLen
}

// This is why writing a NextAt() method would be really nice.
func (gen *sysRandGenerator) Next() float64 {
	if gen.idx == 0 { 
		gen.idx = sysBufLen
		n, err := rand.Read(gen.buf)
		if n != len(gen.buf) {
			panic(fmt.Sprintf("Wrong number of bytes read; %d", n))
		} else if err != nil {
			panic(err.Error)
		}
	}

	gen.idx--
	x, _ := binary.Uvarint(gen.buf[gen.idx * 8: (gen.idx + 1) * 8])
	res := float64(x) / sysMaxUint64
	if res == 1.0 { return gen.Next() }
	return res
}
