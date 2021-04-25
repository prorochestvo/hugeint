package instruction

// Add ...
func Add(a, b uint32) (hDWord, lDWord uint32) {
	return addV2(a, b)
}

func addV1(a, b uint32) (hDWord, lDWord uint32) {
	l, r := a&bit01, b&bit01
	lDWord |= r ^ l
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1 // if (l > 0 && r > 0) || ((l > 0 || r > 0) && hDWord > 0) then hDWord = bit01 << 1

	l, r = a&bit02, b&bit02
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit03, b&bit03
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit04, b&bit04
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit05, b&bit05
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit06, b&bit06
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit07, b&bit07
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit08, b&bit08
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit09, b&bit09
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit10, b&bit10
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit11, b&bit11
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit12, b&bit12
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit13, b&bit13
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit14, b&bit14
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit15, b&bit15
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit16, b&bit16
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit17, b&bit17
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit18, b&bit18
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit19, b&bit19
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit20, b&bit20
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit21, b&bit21
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit22, b&bit22
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit23, b&bit23
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit24, b&bit24
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit25, b&bit25
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit26, b&bit26
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit27, b&bit27
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit28, b&bit28
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit29, b&bit29
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit30, b&bit30
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit31, b&bit31
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) << 1

	l, r = a&bit32, b&bit32
	lDWord |= r ^ l ^ hDWord
	hDWord = ((r & l) | ((r | l) & hDWord)) >> 31

	return
}

func addV2(a, b uint32) (hDWord, lDWord uint32) {
	for bit := uint32(bit01 & 0xFFFFFFFF); bit > 0; bit <<= 1 {
		hDWord <<= 1
		l, r := a&bit, b&bit
		lDWord |= r ^ l ^ hDWord
		hDWord = (r & l) | ((r | l) & hDWord)
	}
	hDWord >>= 31
	return
}
