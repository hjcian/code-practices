package searching

import "testing"

func Test_binarySearch(t *testing.T) {
	nums := []int{17, 20, 26, 31, 44, 54, 55, 77, 93}
	tests := []struct {
		name   string
		target int
		want   int
	}{
		{
			"found 1",
			17, 0,
		},
		{
			"found 2",
			44, 4,
		},
		{
			"not found 1",
			100, -1,
		},
		{
			"not found 2",
			0, -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"-binarySearch", func(t *testing.T) {
			if got := binarySearch(nums, tt.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name+"-interpolationSearch", func(t *testing.T) {
			if got := interpolationSearch(nums, tt.target); got != tt.want {
				t.Errorf("interpolationSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
