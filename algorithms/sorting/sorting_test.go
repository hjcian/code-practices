package sort

import (
	"reflect"
	"testing"
)

type sortFn func(nums []int) []int

func deepCopy(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	return dest
}

func Test_Sortings(t *testing.T) {
	giveNums := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
	wantNums := []int{17, 20, 26, 31, 44, 54, 55, 77, 93}
	tests := []struct {
		name string
		fn   sortFn
	}{
		{"selectionSort", selectionSort},
		{"insertionSort", insertionSort},
		{"insertionSort_20210531", insertionSort_20210531},
		{"insertionSort_20220417", insertionSort_20220417},
		{"bubbleSort", bubbleSort},
		{"mergeSort", mergeSort},
		{"quickSort", quickSort},
		{"heapSort", heapSort},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn(deepCopy(giveNums)); !reflect.DeepEqual(got, wantNums) {
				t.Errorf("got=%v, want=%v", got, wantNums)
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
	b.Run("quickSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
			quickSort(nums)
		}
	})
	b.Run("heapSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
			heapSort(nums)
		}
	})
}
