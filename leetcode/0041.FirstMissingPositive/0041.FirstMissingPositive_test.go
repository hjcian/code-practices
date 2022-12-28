package firstmissingpositive

import "testing"

func Test_0041_FirstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "ex 1",
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			name: "ex 2",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			name: "ex 3",
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
		{
			name: "ex 4",
			nums: []int{3, 4, 1, 1},
			want: 2,
		},
		{
			name: "163 / 173 testcases passed",
			nums: []int{1},
			want: 2,
		},
		{
			name: "164 / 173 testcases passed",
			nums: []int{2, 1},
			want: 3,
		},
		{
			name: "172 / 173 testcases passed",
			nums: []int{0, 3},
			want: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := firstMissingPositive(tt.nums)
			if got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
