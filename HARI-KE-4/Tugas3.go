package main

import (
	"fmt"
	"math"
)

func minMax(numbers ...int) (min, max *int) {
	// Inisialisasi min dan max dengan nilai paling besar dan paling kecil
	min = new(int)
	*min = math.MaxInt32
	max = new(int)
	*max = math.MinInt32

	// Iterasi melalui semua angka inputan
	for _, number := range numbers {
		// Jika angka lebih kecil dari min, maka update min
		if number < *min {
			*min = number
		}

		// Jika angka lebih besar dari max, maka update max
		if number > *max {
			*max = number
		}
	}

	// Return min dan max
	return min, max
}

func main() {
	// Input 6 angka
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80}

	// Panggil fungsi minMax
	min, max := minMax(numbers...)

	// Cetak nilai min dan max
	fmt.Println("Nilai minimum:", *min)
	fmt.Println("Nilai maksimum:", *max)
}
