package allones

func isAllEvenOnes(n int32) bool {
	// Verbose style
	verifier := int32(0b01)
	mask := int32(0b11)
	flag := int32(0b01)
	for i := 0; i < 16; i++ {
		res := (n >> (2 * i)) & mask
		flag &= (res & verifier)
	}
	return verifier&flag > 0
}

func isAllOddOnes(n int32) bool {
	// Ninja style
	flag := int32(0b10)
	for i := 0; i < 16; i++ {
		res := (n >> (2 * i)) & int32(0b11)
		flag &= (res & int32(0b10))
	}
	return int32(0b10)&flag > 0
}
