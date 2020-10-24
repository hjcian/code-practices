package allones

func isAllEventOnes(n int32) bool {
	verifier := int32(0b01)
	mask := int32(0b11)
	flag := int32(0b00)
	expect := int32(0b00)
	for i := 0; i < 16; i++ {
		res := (n >> (2 * i)) & mask
		flag |= (res ^ verifier)
	}
	return expect == flag
}

func isAllOddOnes(n int32) bool {
	flag := int32(0b00)
	for i := 0; i < 16; i++ {
		flag |= ((n>>(2*i))&int32(0b11) ^ int32(0b10))
	}
	return int32(0b00) == flag
}
