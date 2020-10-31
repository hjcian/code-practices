package bloomfilter

import (
	"testing"
)

func Test_bloomFilter_basic(t *testing.T) {
	type args struct {
		add  int
		has  int
		want bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Example 1",
			args{123, 123, true},
		},
		{
			"Example 2",
			args{456, 789, false},
		},
	}

	bf := NewBloomFilter(64)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bf.Add(tt.args.add)
			if got := bf.Has(tt.args.has); got != tt.args.want {
				t.Errorf("[%v] got %v, want %v \n", tt.name, got, tt.args.want)
			}
		})
	}
}
