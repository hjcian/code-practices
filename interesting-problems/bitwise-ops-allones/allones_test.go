package allones

import "testing"

func Test_allOnes(t *testing.T) {
	t.Run("allOddOnes", func(t *testing.T) {
		// t.Skip()
		testCases := []struct {
			n    int32
			want bool
		}{
			{^int32(0b01010101010101010101010101010101), true},
			{^int32(0b00000001010101010101010101010101), true},
			{^int32(0b01110101010101010101010101010101), false},
			{^int32(0b01010101010101010101010101011101), false},
		}
		for _, test := range testCases {
			if got := isAllOddOnes(test.n); got != test.want {
				t.Errorf("[%08b] Want %v, got %v \n", uint8(test.n), test.want, got)
			}
		}
	})
	t.Run("allEvenOnes", func(t *testing.T) {
		// t.Skip()
		testCases := []struct {
			n    int32
			want bool
		}{
			{int32(0b01010101010101010101010101010101), true},
			{int32(0b01110111010101010101010101010101), true},
			{int32(0b00110111010101010101010101010101), false},
			{int32(0b01110111010101010101010101000101), false},
		}
		for _, test := range testCases {
			if got := isAllEvenOnes(test.n); got != test.want {
				t.Errorf("[%08b] Want %v, got %v \n", uint8(test.n), test.want, got)
			}
		}
	})
}
