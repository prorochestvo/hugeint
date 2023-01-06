package benchmarks

import (
	"github.com/prorochestvo/hugeint/internal/instruction"
	"testing"
)

type funcAdd func(uint32, uint32) (uint32, uint32)

func BenchmarkAdd_V1(b *testing.B) {
	benchmarkAdd(b, instruction.AddV1)
}

func BenchmarkAdd_V2(b *testing.B) {
	benchmarkAdd(b, instruction.AddV2)
}

func benchmarkAdd(t *testing.B, f funcAdd) {
	for i := 0; i < t.N; i++ {
		f(uint32(0xFFFFFFFF), uint32(0xFFFFFFFF))
		f(uint32(0x0000FF00), uint32(0x0000F0FF))
		f(uint32(0x0000F00F), uint32(0x0000FFF0))
		f(uint32(0x0000FFFF), uint32(0x000F0FFF))
		f(uint32(0x00000FF0), uint32(0x0000F0FF))
		f(uint32(0x00000001), uint32(0x00000000))
		f(uint32(0x00000000), uint32(0x00000000))
	}
}
