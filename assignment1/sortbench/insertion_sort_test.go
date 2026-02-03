package sortbench

import (
	"math/rand"
	"strconv"
	"testing"
)

func generateRandomSlice(size int) []int {
	r := rand.New(rand.NewSource(42))
	slice := make([]int, size)
	for i := range slice {
		slice[i] = r.Intn(1000)
	}
	return slice
}

// func BenchmarkInsertionSort(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		arr := []int{9, 5, 1, 4, 3, 8}
// 		insertionSort(arr)
// 	}
// }

func BenchmarkInsertionSort_RandomSizes(b *testing.B) {
	sizes := []int{10, 100, 500, 1000, 2000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			data := generateRandomSlice(size)

			b.ResetTimer()
			for b.Loop() {
				arr := make([]int, size)
				copy(arr, data)
				insertionSort(arr)
			}
		})
	}
}
