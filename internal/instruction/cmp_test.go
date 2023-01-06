package instruction

import "testing"

func TestCmpV1(t *testing.T) {
	testCmp(t, CmpV1)
}

func TestCmpV2(t *testing.T) {
	testCmp(t, CmpV2)
}

type funcCmp func(uint32, uint32) int8

func testCmp(t *testing.T, f funcCmp) {
	checkCmp(t, f, 0xFFFFFFFF, 0xFFFFFFFF)
	checkCmp(t, f, 0x00000000, 0x00000001)
	checkCmp(t, f, 0x00000001, 0x00000000)
	checkCmp(t, f, 0x00000003, 0x00000003)
	checkCmp(t, f, 0x00000007, 0x00000007)
	checkCmp(t, f, 0x00000001, 0x00000001)
	checkCmp(t, f, 0x00000003, 0x00000002)
	checkCmp(t, f, 0x000000FF, 0x00000001)
	checkCmp(t, f, 0xFFFFFFFF, 0x00000001)
	checkCmp(t, f, 0x00000000, 0x00000000)
	checkCmp(t, f, 0x0F000000, 0xE0000000)
	checkCmp(t, f, 0x0F000000, 0x00F00000)

	for i := uint64(0); i <= 0xFF; i++ {
		for j := uint64(0); j <= 0xFF; j++ {
			if !checkCmp(t, f, i, j) {
				return
			}
		}
	}
}

func checkCmp(t *testing.T, f funcCmp, a uint64, b uint64) bool {
	r := f(uint32(a&0xFFFFFFFF), uint32(b&0xFFFFFFFF))
	if v := r < 0; v != (a < b) {
		t.Errorf("comparison (%d < %d) result is not correctly; expected: %t; actual: %t; result = %d;", a, b, a < b, v, r)
		return false
	}
	if v := r > 0; v != (a > b) {
		t.Errorf("comparison (%d > %d) result is not correctly; expected: %t; actual: %t; result = %d;", a, b, a > b, v, r)
		return false
	}
	if v := r == 0; v != (a == b) {
		t.Errorf("comparison (%d == %d) result is not correctly; expected: %t; actual: %t; result = %d;", a, b, a == b, v, r)
		return false
	}
	return true
}
