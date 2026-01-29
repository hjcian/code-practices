package topkfrequenteventswithorderpreservation

import "testing"

func TestGetTopKFrequentEvents(t *testing.T) {
	tests := []struct {
		name   string
		events []int32
		k      int32
		want   []int32
	}{
		{
			name:   "basic case",
			events: []int32{1, 1, 1, 2, 2, 3},
			k:      2,
			want:   []int32{1, 2},
		},
		{
			name:   "order preservation",
			events: []int32{1, 2, 3, 1, 2, 3, 1},
			k:      2,
			want:   []int32{1, 2},
		},
		{
			name:   "single element",
			events: []int32{5, 5, 5},
			k:      1,
			want:   []int32{5},
		},
		{
			name:   "all same frequency",
			events: []int32{1, 2, 3},
			k:      2,
			want:   []int32{1, 2},
		},
		{
			name:   "k equals array length",
			events: []int32{1, 2, 1, 3},
			k:      3,
			want:   []int32{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTopKFrequentEvents(tt.events, tt.k)
			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
