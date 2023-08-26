package main

import "fmt"

func Mapping(slice []string) map[string]int {
	// Membuat map yang akan menyimpan informasi berapa kali setiap string muncul
	stringCount := make(map[string]int)

	// Mengiterasi setiap string dalam slice
	for _, s := range slice {
		// Menambahkan jumlah kemunculan string ke dalam map
		// Jika string sudah ada dalam map, kita tambahkan 1 pada jumlahnya
		// Jika string belum ada dalam map, kita inisialisasi jumlahnya dengan 1
		stringCount[s]++
	}

	return stringCount
}

func main() {
	// Test cases
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	// Output: map[adi:1 asd:2 qwe:3]

	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	// Output: map[asd:2 qwe:1]

	fmt.Println(Mapping([]string{}))
	// Output: map[]
}
