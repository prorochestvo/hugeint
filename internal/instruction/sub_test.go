package instruction

import (
	"fmt"
	"testing"
)

func TestSubV1(t *testing.T) {
	testSub(t, SubV1)
}

type funcSub func(uint32, uint32) (uint32, uint32)

func testSub(t *testing.T, f funcSub) {
	if err := checkSub(f, 0x00000006, 0x00000005); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000005, 0x00000007); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000000, 0x00000001); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000000, 0x00000002); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000000, 0x00000003); err != nil {
		t.Error(err)
	}

	if err := checkSub(f, 0x00000003, 0x00000003); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000002, 0x00000001); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000002, 0x00000000); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0xFFFFFFFF, 0xFFFFFFFF); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000001, 0x00000000); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000007, 0x00000007); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000001, 0x00000001); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000003, 0x00000002); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x000000FF, 0x00000001); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0xFFFFFFFF, 0x00000001); err != nil {
		t.Error(err)
	}
	if err := checkSub(f, 0x00000000, 0x00000000); err != nil {
		t.Error(err)
	}

	for i := uint64(0); i <= 0x0F; i++ {
		for j := uint64(0); j <= 0x0F; j++ {
			if err := checkSub(f, j, i); err != nil {
				t.Error(err)
			}
		}
	}
}

func checkSub(f funcSub, a uint64, b uint64) error {
	hDW, lDW := f(uint32(a&0xFFFFFFFF), uint32(b&0xFFFFFFFF))
	if v := (uint64(hDW) << 32) | uint64(lDW); int64(v) != (int64(a) - int64(b)) {
		return fmt.Errorf("subtract (%d - %d) result is not correctly; expected: %b; actual: %b; result = H:%b | L:%b", a, b, int64(a)-int64(b), int64(v), hDW, lDW)
	}
	return nil
}
