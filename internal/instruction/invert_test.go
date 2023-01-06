package instruction

import "testing"

func TestInvertV1(t *testing.T) {
	testInvert(t, invertV1)
}

func TestInvertV2(t *testing.T) {
	testInvert(t, invertV2)
}

type funcInvert func(uint32) uint32

func testInvert(t *testing.T, f funcInvert) {
	checkInvert(t, f, 0xFFFFFFFF, 0x00000000)
	checkInvert(t, f, 0x00000000, 0xFFFFFFFF)
	checkInvert(t, f, 0x00000001, 0xFFFFFFFE)
	checkInvert(t, f, 0x00000003, 0xFFFFFFFC)
	checkInvert(t, f, 0x00000007, 0xFFFFFFF8)
	checkInvert(t, f, 0xFF000001, 0x00FFFFFE)
}

func checkInvert(t *testing.T, f funcInvert, v uint32, expectingResult uint32) bool {
	res := f(v & 0xFFFFFFFF)
	if res != expectingResult {
		t.Errorf("addition result is not correctly; expected: %d (%d); actual: %d;", expectingResult, v, res)
		return false
	}
	return true
}
