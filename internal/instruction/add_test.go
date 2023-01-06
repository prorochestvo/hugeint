package instruction

import (
	"fmt"
	"testing"
)

func TestAddV1(t *testing.T) {
	testAdd(t, AddV1)
}

func TestAddV2(t *testing.T) {
	testAdd(t, AddV2)
}

type funcAdd func(uint32, uint32) (uint32, uint32)

func testAdd(t *testing.T, f funcAdd) {
	if err := checkAdd(f, 0xFFFFFFFF, 0xFFFFFFFF); err != nil {
		t.Error(err)
	}
	if err := checkAdd(f, 0x00000000, 0x00000001); err != nil {
		t.Error(err)
	}
	if err := checkAdd(f, 0x00000001, 0x00000000); err != nil {
		t.Error(err)
	}
	if err := checkAdd(f, 0x00000003, 0x00000003); err != nil {
		t.Error(err)
	}
	if err := checkAdd(f, 0x00000007, 0x00000007); err != nil {
		t.Error(err)
	}
	if err := checkAdd(f, 0x00000001, 0x00000001); err != nil {
		t.Error(err)
	}
	if err := checkAdd(f, 0x00000003, 0x00000002); err != nil {
		t.Error(err)
	}
	if err := checkAdd(f, 0x000000FF, 0x00000001); err != nil {
		t.Error(err)
	}
	if err := checkAdd(f, 0xFFFFFFFF, 0x00000001); err != nil {
		t.Error(err)
	}
	if err := checkAdd(f, 0x00000000, 0x00000000); err != nil {
		t.Error(err)
	}

	for i := uint64(0); i <= 0xFF; i++ {
		for j := uint64(0); j <= 0xFF; j++ {
			if err := checkAdd(f, i, j); err != nil {
				t.Error(err)
			}
		}
	}
}

func checkAdd(f funcAdd, a uint64, b uint64) error {
	hDW, lDW := f(uint32(a&0xFFFFFFFF), uint32(b&0xFFFFFFFF))
	if v := (uint64(hDW) << 32) | uint64(lDW); v != a+b {
		return fmt.Errorf("addition (%d + %d) result is not correctly; expected: %b; actual: %b; result = H:%b | L:%b", a, b, a+b, v, hDW, lDW)
	}
	return nil
}
