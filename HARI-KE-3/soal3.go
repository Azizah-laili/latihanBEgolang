package main

import (
	"fmt"
)

func MunculSekali(angka string) []int {
	// Membuat map untuk menyimpan jumlah kemunculan setiap angka
	angkaCount := make(map[int]int)

	// Menghitung kemunculan setiap angka dalam input
	for _, digit := range angka {
		angkaCount[int(digit)-'0']++
	}

	// Membuat slice untuk menyimpan angka yang muncul sekali
	uniqueDigits := []int{}
	for digit, count := range angkaCount {
		if count == 1 {
			uniqueDigits = append(uniqueDigits, digit)
		}
	}

	return uniqueDigits
}

func main() {
	fmt.Println(MunculSekali("1234123"))    // Output: [4]
	fmt.Println(MunculSekali("76523752"))   // Output: [6 3]
	fmt.Println(MunculSekali("12345"))      // Output: [1 2 3 4 5]
	fmt.Println(MunculSekali("1122334455")) // Output: []
	fmt.Println(MunculSekali("0872504"))    // Output: [8 7 2 5 4]
}
