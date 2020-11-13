package interlacedspiralcipher

import (
	"fmt"
	"math"
	"strings"
)

func displayMatrix(m [][]byte) {
	for i := range m {
		fmt.Println(m[i])
	}
}

func getCharOrPadSpace(s string, i int) byte {
	if i >= len(s) {
		return []byte(" ")[0]
	}
	return s[i]
}

type cipherFunc func(s string) string

var encode cipherFunc = func(s string) string {
	// your code goes here. you can do it!
	sideLen := int(math.Ceil(math.Sqrt(float64(len(s)))))
	matrix := make([][]byte, sideLen, sideLen)
	for i := range matrix {
		matrix[i] = make([]byte, sideLen, sideLen)
	}

	sIdx := 0
	for begin, end := 0, sideLen-1; begin <= end; begin, end = begin+1, end-1 {
		for i := 0; i < end-begin; i++ {
			matrix[begin][begin+i] = getCharOrPadSpace(s, sIdx)
			sIdx++
			matrix[begin+i][end] = getCharOrPadSpace(s, sIdx)
			sIdx++
			matrix[end][end-i] = getCharOrPadSpace(s, sIdx)
			sIdx++
			matrix[end-i][begin] = getCharOrPadSpace(s, sIdx)
			sIdx++
		}
		if begin == end {
			// handle the center of matrix case
			matrix[begin][begin] = getCharOrPadSpace(s, sIdx)
		}
	}
	// concate strings
	ret := ""
	for i := range matrix {
		ret += string(matrix[i])
	}
	return ret
}

var decode cipherFunc = func(s string) string {
	// your code goes here. you can do it!
	sideLen := int(math.Ceil(math.Sqrt(float64(len(s)))))
	matrix := make([][]byte, sideLen, sideLen)
	for i := range matrix {
		matrix[i] = make([]byte, sideLen, sideLen)
		for j := range matrix[i] {
			// fill the characters into matrix
			matrix[i][j] = s[i*sideLen+j]
		}
	}

	decoded := make([]byte, len(s), len(s))
	sIdx := 0
	for begin, end := 0, sideLen-1; begin <= end; begin, end = begin+1, end-1 {
		for i := 0; i < end-begin; i++ {
			decoded[sIdx] = matrix[begin][begin+i]
			sIdx++
			decoded[sIdx] = matrix[begin+i][end]
			sIdx++
			decoded[sIdx] = matrix[end][end-i]
			sIdx++
			decoded[sIdx] = matrix[end-i][begin]
			sIdx++
		}
		if begin == end {
			// handle the center of matrix case
			decoded[sIdx] = matrix[begin][begin]
		}
	}
	return strings.TrimSpace(string(decoded))
}

var InterlacedSpiralCipher = map[string]cipherFunc{"encode": encode, "decode": decode}
