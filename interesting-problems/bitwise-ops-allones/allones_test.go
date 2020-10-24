package allones

import "testing"

func Test_allEvenOnes(t *testing.T) {
	t.Run("allOddOnes", func(t *testing.T) {
		// t.Skip()
		testCases := []struct {
			n    int32
			want bool
		}{
			{^int32(0b01010101010101010101010101010101), true},
			{^int32(0b01010101010101010101010111010101), false},
			{^int32(0b01010101011101110101010111010101), false},
			{int32(0b01010101011101110101010111010101), false},
			{int32(0b01110101011101110101010111010101), false},
			{int32(0b01110111011101110101010111010101), false},
			{int32(0b01110100011101110101010111010101), false},
		}
		for _, test := range testCases {
			if got := isAllOddOnes(test.n); got != test.want {
				t.Errorf("[%08b] Want %v, got %v \n", uint8(test.n), test.want, got)
			}
		}
	})
	t.Run("allEventOnes", func(t *testing.T) {
		testCases := []struct {
			n    int32
			want bool
		}{
			{int32(0b01010101010101010101010101010101), true},
			{int32(0b01010101010101010101010101000101), false},
			{^int32(0b01010101010101010101010101010101), false},
			{int32(0b01010101010101010101010111010101), false},
			{int32(0b01010101011101110101010111010101), false},
		}
		for _, test := range testCases {
			if got := isAllEventOnes(test.n); got != test.want {
				t.Errorf("[%08b] Want %v, got %v \n", uint8(test.n), test.want, got)
			}
		}
	})
}
