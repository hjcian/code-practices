package findfirstoccurrence

import "testing"

func TestFindFirstOccurrence(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int32
		target int32
		want   int32
	}{
		{"target at beginning", []int32{1, 2, 3, 4, 5}, 1, 0},
		{"target at end", []int32{1, 2, 3, 4, 5}, 5, 4},
		{"target in middle", []int32{1, 2, 3, 4, 5}, 3, 2},
		{"target not found", []int32{1, 2, 3, 4, 5}, 6, -1},
		{"duplicate targets", []int32{1, 2, 2, 2, 3}, 2, 1},
		{"single element match", []int32{5}, 5, 0},
		{"single element no match", []int32{5}, 3, -1},
		{"empty array", []int32{}, 1, -1},
		{"all same elements", []int32{2, 2, 2, 2}, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFirstOccurrence(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("findFirstOccurrence(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
