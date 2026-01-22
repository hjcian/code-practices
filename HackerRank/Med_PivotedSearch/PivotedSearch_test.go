package pivotedsearch

import "testing"

func TestSearchRotatedTimestamps(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int32
		target int32
		want   int32
	}{
		{
			name:   "target found in left portion",
			nums:   []int32{4, 5, 6, 7, 0, 1, 2},
			target: 5,
			want:   1,
		},
		{
			name:   "target found in right portion",
			nums:   []int32{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "target not found",
			nums:   []int32{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "single element found",
			nums:   []int32{1},
			target: 1,
			want:   0,
		},
		{
			name:   "single element not found",
			nums:   []int32{1},
			target: 2,
			want:   -1,
		},
		{
			name:   "target at start",
			nums:   []int32{4, 5, 6, 7, 0, 1, 2},
			target: 4,
			want:   0,
		},
		{
			name:   "target at end",
			nums:   []int32{4, 5, 6, 7, 0, 1, 2},
			target: 2,
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRotatedTimestamps(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("searchRotatedTimestamps(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
