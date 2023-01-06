package instruction

// http://reshinfo.com/primer_vichitanije2.php

func SubV1(a, b uint32) (hDWord, lDWord uint32) {
	hDWord = 0
	lDWord = 0
	if cmp := CmpV2(a, b); cmp == 0 {
		return
	} else if cmp < 0 {
		a = a ^ 0xFFFFFFFF
		if a == 0xFFFFFFFF {
			a = 0
			//a1 := 0x00000001
		} else {
			a += 1 // TODO: rethink
		}
		hDWord, lDWord = AddV2(a, b)
		//lDWord = lDWord ^ 0xFFFFFFFF
		//hDWord = hDWord ^ 0xFFFFFFFF
		hDWord = 0
		//lDWord -= 1
		//if hDWord != 0 {
		//hDWord -= 1
		//} else {
		//	lDWord -= 1
		//}
	} else {
		l, r := a&bit01, b&bit01
		lDWord |= r ^ l
		hDWord = (((l ^ bit01) & r) | ((l ^ bit01) & (r ^ bit01) & hDWord)) << 1

		l, r = a&bit02, b&bit02
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit02) & r) | ((l ^ bit02) & (r ^ bit02) & hDWord)) << 1

		l, r = a&bit03, b&bit03
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit03) & r) | ((l ^ bit03) & (r ^ bit03) & hDWord)) << 1

		l, r = a&bit04, b&bit04
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit04) & r) | ((l ^ bit04) & (r ^ bit04) & hDWord)) << 1

		l, r = a&bit05, b&bit05
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit05) & r) | ((l ^ bit05) & (r ^ bit05) & hDWord)) << 1

		l, r = a&bit06, b&bit06
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit06) & r) | ((l ^ bit06) & (r ^ bit06) & hDWord)) << 1

		l, r = a&bit07, b&bit07
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit07) & r) | ((l ^ bit07) & (r ^ bit07) & hDWord)) << 1

		l, r = a&bit08, b&bit08
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit08) & r) | ((l ^ bit08) & (r ^ bit08) & hDWord)) << 1

		l, r = a&bit09, b&bit09
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit09) & r) | ((l ^ bit09) & (r ^ bit09) & hDWord)) << 1

		l, r = a&bit10, b&bit10
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit10) & r) | ((l ^ bit10) & (r ^ bit10) & hDWord)) << 1

		l, r = a&bit11, b&bit11
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit11) & r) | ((l ^ bit11) & (r ^ bit11) & hDWord)) << 1

		l, r = a&bit12, b&bit12
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit12) & r) | ((l ^ bit12) & (r ^ bit12) & hDWord)) << 1

		l, r = a&bit13, b&bit13
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit13) & r) | ((l ^ bit13) & (r ^ bit13) & hDWord)) << 1

		l, r = a&bit14, b&bit14
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit14) & r) | ((l ^ bit14) & (r ^ bit14) & hDWord)) << 1

		l, r = a&bit15, b&bit15
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit15) & r) | ((l ^ bit15) & (r ^ bit15) & hDWord)) << 1

		l, r = a&bit16, b&bit16
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit16) & r) | ((l ^ bit16) & (r ^ bit16) & hDWord)) << 1

		l, r = a&bit17, b&bit17
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit17) & r) | ((l ^ bit17) & (r ^ bit17) & hDWord)) << 1

		l, r = a&bit18, b&bit18
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit18) & r) | ((l ^ bit18) & (r ^ bit18) & hDWord)) << 1

		l, r = a&bit19, b&bit19
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit19) & r) | ((l ^ bit19) & (r ^ bit19) & hDWord)) << 1

		l, r = a&bit20, b&bit20
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit20) & r) | ((l ^ bit20) & (r ^ bit20) & hDWord)) << 1

		l, r = a&bit21, b&bit21
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit21) & r) | ((l ^ bit21) & (r ^ bit21) & hDWord)) << 1

		l, r = a&bit22, b&bit22
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit22) & r) | ((l ^ bit22) & (r ^ bit22) & hDWord)) << 1

		l, r = a&bit23, b&bit23
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit23) & r) | ((l ^ bit23) & (r ^ bit23) & hDWord)) << 1

		l, r = a&bit24, b&bit24
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit24) & r) | ((l ^ bit24) & (r ^ bit24) & hDWord)) << 1

		l, r = a&bit25, b&bit25
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit25) & r) | ((l ^ bit25) & (r ^ bit25) & hDWord)) << 1

		l, r = a&bit26, b&bit26
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit26) & r) | ((l ^ bit26) & (r ^ bit26) & hDWord)) << 1

		l, r = a&bit27, b&bit27
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit27) & r) | ((l ^ bit27) & (r ^ bit27) & hDWord)) << 1

		l, r = a&bit28, b&bit28
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit28) & r) | ((l ^ bit28) & (r ^ bit28) & hDWord)) << 1

		l, r = a&bit29, b&bit29
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit29) & r) | ((l ^ bit29) & (r ^ bit29) & hDWord)) << 1

		l, r = a&bit30, b&bit30
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit30) & r) | ((l ^ bit30) & (r ^ bit30) & hDWord)) << 1

		l, r = a&bit31, b&bit31
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit31) & r) | ((l ^ bit31) & (r ^ bit31) & hDWord)) << 1

		l, r = a&bit32, b&bit32
		lDWord |= r ^ l ^ hDWord
		hDWord = (((l ^ bit32) & r) | ((l ^ bit32) & (r ^ bit32) & hDWord)) >> 31
	}

	return
}
