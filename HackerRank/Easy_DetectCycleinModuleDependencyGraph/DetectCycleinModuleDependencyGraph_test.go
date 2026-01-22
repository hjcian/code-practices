package detectcycleinmoduledependencygraph

import "testing"

func TestHasCircularDependency(t *testing.T) {
	tests := []struct {
		name         string
		n            int32
		dependencies [][]int32
		expected     bool
	}{
		{
			name:         "no cycle",
			n:            3,
			dependencies: [][]int32{{1, 2}, {2, 3}},
			expected:     false,
		},
		{
			name:         "simple cycle",
			n:            3,
			dependencies: [][]int32{{1, 2}, {2, 3}, {3, 1}},
			expected:     true,
		},
		{
			name:         "self dependency",
			n:            1,
			dependencies: [][]int32{{1, 1}},
			expected:     true,
		},
		{
			name:         "empty dependencies",
			n:            3,
			dependencies: [][]int32{},
			expected:     false,
		},
		{
			name:         "two node cycle",
			n:            2,
			dependencies: [][]int32{{1, 2}, {2, 1}},
			expected:     true,
		},
		{
			name:         "complex no cycle",
			n:            5,
			dependencies: [][]int32{{1, 2}, {1, 3}, {2, 4}, {3, 4}, {4, 5}},
			expected:     false,
		},
		{
			name:         "complex cycle",
			n:            5,
			dependencies: [][]int32{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {4, 5}},
			expected:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasCircularDependency(tt.n, tt.dependencies)
			if result != tt.expected {
				t.Errorf("hasCircularDependency(%d, %v) = %v, want %v", tt.n, tt.dependencies, result, tt.expected)
			}
		})
	}
}
