package instruction

func CmpV1(a, b uint32) int8 {
	for bit := bit32 & 0xFFFFFFFF; bit > 0; bit >>= 1 {
		l, r := a&bit, b&bit
		if l < r {
			return -32
		}
		if l > r {
			return +32
		}
	}
	return 0
}

func CmpV2(a, b uint32) int8 {
	l, r := a&bit32, b&bit32
	if l < r {
		return -32
	}
	if l > r {
		return +32
	}

	l, r = a&bit31, b&bit31
	if l < r {
		return -31
	}
	if l > r {
		return +31
	}

	l, r = a&bit30, b&bit30
	if l < r {
		return -30
	}
	if l > r {
		return +30
	}

	l, r = a&bit29, b&bit29
	if l < r {
		return -29
	}
	if l > r {
		return +29
	}

	l, r = a&bit28, b&bit28
	if l < r {
		return -28
	}
	if l > r {
		return +28
	}

	l, r = a&bit27, b&bit27
	if l < r {
		return -27
	}
	if l > r {
		return +27
	}

	l, r = a&bit26, b&bit26
	if l < r {
		return -26
	}
	if l > r {
		return +26
	}

	l, r = a&bit25, b&bit25
	if l < r {
		return -25
	}
	if l > r {
		return +25
	}

	l, r = a&bit24, b&bit24
	if l < r {
		return -24
	}
	if l > r {
		return +24
	}

	l, r = a&bit23, b&bit23
	if l < r {
		return -23
	}
	if l > r {
		return +23
	}

	l, r = a&bit22, b&bit22
	if l < r {
		return -22
	}
	if l > r {
		return +22
	}

	l, r = a&bit21, b&bit21
	if l < r {
		return -21
	}
	if l > r {
		return +21
	}

	l, r = a&bit20, b&bit20
	if l < r {
		return -20
	}
	if l > r {
		return +20
	}

	l, r = a&bit19, b&bit19
	if l < r {
		return -19
	}
	if l > r {
		return +19
	}

	l, r = a&bit18, b&bit18
	if l < r {
		return -18
	}
	if l > r {
		return +18
	}

	l, r = a&bit17, b&bit17
	if l < r {
		return -17
	}
	if l > r {
		return +17
	}

	l, r = a&bit16, b&bit16
	if l < r {
		return -16
	}
	if l > r {
		return +16
	}

	l, r = a&bit15, b&bit15
	if l < r {
		return -15
	}
	if l > r {
		return +15
	}

	l, r = a&bit14, b&bit14
	if l < r {
		return -14
	}
	if l > r {
		return +14
	}

	l, r = a&bit13, b&bit13
	if l < r {
		return -13
	}
	if l > r {
		return +13
	}

	l, r = a&bit12, b&bit12
	if l < r {
		return -12
	}
	if l > r {
		return +12
	}

	l, r = a&bit11, b&bit11
	if l < r {
		return -11
	}
	if l > r {
		return +11
	}

	l, r = a&bit10, b&bit10
	if l < r {
		return -10
	}
	if l > r {
		return +10
	}

	l, r = a&bit09, b&bit09
	if l < r {
		return -9
	}
	if l > r {
		return +9
	}

	l, r = a&bit08, b&bit08
	if l < r {
		return -8
	}
	if l > r {
		return +8
	}

	l, r = a&bit07, b&bit07
	if l < r {
		return -7
	}
	if l > r {
		return +7
	}

	l, r = a&bit06, b&bit06
	if l < r {
		return -6
	}
	if l > r {
		return +6
	}

	l, r = a&bit05, b&bit05
	if l < r {
		return -5
	}
	if l > r {
		return +5
	}

	l, r = a&bit04, b&bit04
	if l < r {
		return -4
	}
	if l > r {
		return +4
	}

	l, r = a&bit03, b&bit03
	if l < r {
		return -3
	}
	if l > r {
		return +3
	}

	l, r = a&bit02, b&bit02
	if l < r {
		return -2
	}
	if l > r {
		return +2
	}

	l, r = a&bit01, b&bit01
	if l < r {
		return -1
	}
	if l > r {
		return +1
	}

	return 0
}
