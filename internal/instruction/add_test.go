package instruction

import "testing"

type funcAdd func(uint32, uint32) (uint32, uint32)

func TestAdd(t *testing.T) {
	testAdd(t, Add)
}

func TestAddV1(t *testing.T) {
	testAdd(t, addV1)
}

func TestAddV2(t *testing.T) {
	testAdd(t, addV2)
}

func BenchmarkAdd_V1_FFFFFFFF_and_FFFFFFFF(b *testing.B) {
	benchmarkAdd(b, addV1, 0xFFFFFFFF, 0xFFFFFFFF)
}

func BenchmarkAdd_V2_FFFFFFFF_and_FFFFFFFF(b *testing.B) {
	benchmarkAdd(b, addV2, 0xFFFFFFFF, 0xFFFFFFFF)
}

func testAdd(t *testing.T, f funcAdd) {
	checkAdd(t, f, 0xFFFFFFFF, 0xFFFFFFFF)
	checkAdd(t, f, 0x00000000, 0x00000001)
	checkAdd(t, f, 0x00000001, 0x00000000)
	checkAdd(t, f, 0x00000003, 0x00000003)
	checkAdd(t, f, 0x00000007, 0x00000007)
	checkAdd(t, f, 0x00000001, 0x00000001)
	checkAdd(t, f, 0x00000003, 0x00000002)
	checkAdd(t, f, 0x000000FF, 0x00000001)
	checkAdd(t, f, 0xFFFFFFFF, 0x00000001)
	checkAdd(t, f, 0x00000000, 0x00000000)

	for i := uint64(0); i <= 0x0F; i++ {
		for j := uint64(0); j <= 0x0F; j++ {
			if !checkAdd(t, Add, i, j) {
				return
			}
		}
	}
}

func checkAdd(t *testing.T, f funcAdd, a uint64, b uint64) bool {
	hDW, lDW := f(uint32(a&0xFFFFFFFF), uint32(b&0xFFFFFFFF))
	if v := (uint64(hDW) << 32) | uint64(lDW); v != a+b {
		t.Errorf("addition result is not correctly; expected: %d (%d + %d); actual: %d [H:%d|L:%d];", a+b, a, b, v, hDW, lDW)
		return false
	}
	return true
}

func benchmarkAdd(t *testing.B, f funcAdd, a uint64, b uint64) {
	for i := 0; i < t.N; i++ {
		f(uint32(a&0xFFFFFFFF), uint32(b&0xFFFFFFFF))
	}
}
