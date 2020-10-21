package allones

func isAllEventOnes(n int32) bool {
	return (n ^ int32(0b01010101010101010101010101010101)) == 0
}

func isAllOddOnes(n int32) bool {
	return (n ^ ^int32(0b01010101010101010101010101010101)) == 0
}
