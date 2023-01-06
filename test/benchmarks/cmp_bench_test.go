package benchmarks

import (
	"github.com/prorochestvo/hugeint/internal/instruction"
	"testing"
)

type funcCmp func(uint32, uint32) int8

func BenchmarkCmp_V1(b *testing.B) {
	benchmarkCmp(b, instruction.CmpV1)
}

func BenchmarkCmp_V2(b *testing.B) {
	benchmarkCmp(b, instruction.CmpV2)
}

func benchmarkCmp(t *testing.B, f funcCmp) {
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
