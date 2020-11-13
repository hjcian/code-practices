package interlacedspiralcipher

import "testing"

func Test_InterlacedSpiralCipher(t *testing.T) {
	phrase1A := "Romani ite domum"
	phrase1B := "Rntodomiimuea  m"
	phrase2A := "Sic transit gloria mundi"
	phrase2B := "Stsgiriuar i ninmd l otac"
	testCases := []struct {
		dencode string
		input   string
		sol     string
	}{
		{"encode", phrase1A, phrase1B},
		{"decode", phrase1B, phrase1A},
		{"encode", phrase2A, phrase2B},
		{"decode", phrase2B, phrase2A},
	}
	for _, test := range testCases {
		if got := InterlacedSpiralCipher[test.dencode](test.input); got != test.sol {
			t.Errorf("[%v] Expect '%v', got '%v' ", test.dencode, test.sol, got)
		}
	}
}
