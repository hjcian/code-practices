package minimumplanstoreachtargetbandwidth

import "testing"

func TestFindMinimumPlansForBandwidth(t *testing.T) {
	tests := []struct {
		name            string
		planSizes       []int32
		targetBandwidth int32
		expectedResult  int32
	}{
		{
			"no plans available",
			[]int32{},
			10,
			-1,
		},
		{
			name:            "target bandwidth is zero",
			planSizes:       []int32{1, 2, 3},
			targetBandwidth: 0,
			expectedResult:  0,
		},
		{
			name:            "single plan matches target",
			planSizes:       []int32{5},
			targetBandwidth: 5,
			expectedResult:  1,
		},
		{
			name:            "combination of plans",
			planSizes:       []int32{1, 2, 5},
			targetBandwidth: 5,
			expectedResult:  1,
		},
		{
			name:            "multiple plans needed",
			planSizes:       []int32{2, 3},
			targetBandwidth: 7,
			expectedResult:  3,
		},
		{
			name:            "impossible to reach target",
			planSizes:       []int32{2, 4},
			targetBandwidth: 3,
			expectedResult:  -1,
		},
		{
			name:            "large target with small plans",
			planSizes:       []int32{1, 3},
			targetBandwidth: 4,
			expectedResult:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMinimumPlansForBandwidth(tt.planSizes, tt.targetBandwidth)
			if result != tt.expectedResult {
				t.Errorf("got %d, want %d", result, tt.expectedResult)
			}
		})
	}
}
