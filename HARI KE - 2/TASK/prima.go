package main

import "fmt"

// func main() {
// 	for i := 7; i <= 100; i += 7 {
// 		fmt.Println(i)
// 	}
// }

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	number := 17

	if isPrime(number) {
		fmt.Printf("adalah bilanga prima: %d", number)
	} else {
		fmt.Printf("Bukan bilang prima", number)
	}
}
