package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// Buat array kosong untuk menampung data gabungan
	result := []string{}

	// Iterasi array A
	for _, value := range arrayA {
		// Cek apakah value sudah ada di array result
		if !InArray(value, result) {
			// Jika belum ada, tambahkan ke array result
			result = append(result, value)
		}
	}

	// Iterasi array B
	for _, value := range arrayB {
		// Cek apakah value sudah ada di array result
		if !InArray(value, result) {
			// Jika belum ada, tambahkan ke array result
			result = append(result, value)
		}
	}

	return result
}

func InArray(value string, array []string) bool {
	// Iterasi array
	for _, item := range array {
		// Jika value ditemukan, kembalikan true
		if value == item {
			return true
		}
	}

	// Jika value tidak ditemukan, kembalikan false
	return false
}

func main() {
	//test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	//{"king", "devil jin", "akuma", eddie", "steve",  "geese }

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	//{"sergei", "jin", "steve", "bryan"}

	fmt.Println(ArrayMerge([]string{"alisa", "yohimitsu"}, []string{"devil jin", "yohimitsu", "alisa", "law"}))
	//{"devil jin", "yohimitsu", "alisa", "law" }

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "segei"}))
	//["devil jin", "sergei" ]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	//{"hwoarang"}

	fmt.Println(ArrayMerge([]string{}, []string{}))
	//{}
}
