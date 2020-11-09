package pickpeaks

import (
	"reflect"
	"testing"
)

func TestPickPeaks(t *testing.T) {

	tests := []struct {
		name  string
		array []int
		want  PosPeaks
	}{
		{
			"should support finding peaks",
			[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		},
		{
			"should support finding peaks, but should ignore peaks on the edge of the array",
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		},
		{
			"should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau",
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
			PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		},
		{
			"should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau",
			[]int{2, 1, 3, 1, 2, 2, 2, 2, 1},
			PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}},
		},
		{
			"should support finding peaks, but should ignore peaks on the edge of the array",
			[]int{2, 1, 3, 1, 2, 2, 2, 2},
			PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PickPeaks(tt.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PickPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}
