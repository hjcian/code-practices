package targetindexsearch

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int32
		target int32
		want   int32
	}{
		{"target found at start", []int32{1, 3, 5, 7, 9}, 1, 0},
		{"target found in middle", []int32{1, 3, 5, 7, 9}, 5, 2},
		{"target found at end", []int32{1, 3, 5, 7, 9}, 9, 4},
		{"target not found", []int32{1, 3, 5, 7, 9}, 4, -1},
		{"single element found", []int32{5}, 5, 0},
		{"single element not found", []int32{5}, 3, -1},
		{"empty array", []int32{}, 5, -1},
		{"target smaller than all", []int32{2, 4, 6, 8}, 1, -1},
		{"target larger than all", []int32{2, 4, 6, 8}, 10, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("binarySearch(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
