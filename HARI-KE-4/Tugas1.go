package main

import (
	"fmt"
)

type Car struct {
	typeCar string
	fuelIn  float64
}

func (c *Car) calculateDistance() float64 {
	// Menghitung perkiraan jarak yang bisa ditempuh
	// berdasarkan bahan bakar yang sedang terisi
	distance := c.fuelIn * 1.5
	return distance
}

func main() {
	// Membuat objek Car
	car := Car{
		typeCar: "Honda Jazz",
		fuelIn:  10,
	}

	// Menampilkan perkiraan jarak yang bisa ditempuh
	fmt.Println(car.calculateDistance())
}
