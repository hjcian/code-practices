package arrayproblems

import (
	"reflect"
	"testing"
)

func deepCopy(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	return dest
}

func Test_selectionSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Sample",
			args{[]int{54, 26, 93, 17, 77, 31, 44, 55, 20}},
			[]int{17, 20, 26, 31, 44, 54, 55, 77, 93},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(deepCopy(tt.args.nums)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSort(deepCopy(tt.args.nums)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSort() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := bubbleSort(deepCopy(tt.args.nums)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(deepCopy(tt.args.nums)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_all(b *testing.B) {
	b.Run("selectionSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
			selectionSort(nums)
		}
	})
	b.Run("insertionSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
			insertionSort(nums)
		}
	})
	b.Run("bubbleSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
			bubbleSort(nums)
		}
	})
	b.Run("mergeSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
			mergeSort(nums)
		}
	})
}
