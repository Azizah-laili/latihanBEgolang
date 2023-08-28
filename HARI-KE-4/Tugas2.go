package main

import (
	"fmt"
)

type Student struct {
	name  string
	score []int
}

func main() {
	// Buat array kosong untuk menyimpan data siswa
	students := make([]Student, 5)

	// Masukkan data siswa
	students[0] = Student{
		name:  "Rizky",
		score: []int{100, 90, 80},
	}
	students[1] = Student{
		name:  "Kobar",
		score: []int{85, 75, 65},
	}
	students[2] = Student{
		name:  "Ismail",
		score: []int{70, 60, 50},
	}
	students[3] = Student{
		name:  "Umam",
		score: []int{65, 55, 45},
	}
	students[4] = Student{
		name:  "Yopan",
		score: []int{50, 40, 30},
	}

	// Hitung skor rata-rata
	averageScore := 0
	for _, student := range students {
		for _, score := range student.score {
			averageScore += score
		}
	}
	averageScore /= len(students)

	// Cari siswa dengan skor minimum
	minScore := students[0].score[0]
	minStudent := students[0]
	for _, student := range students {
		for _, score := range student.score {
			if score < minScore {
				minScore = score
				minStudent = student
			}
		}
	}

	// Cari siswa dengan skor maksimum
	maxScore := students[0].score[0]
	maxStudent := students[0]
	for _, student := range students {
		for _, score := range student.score {
			if score > maxScore {
				maxScore = score
				maxStudent = student
			}
		}
	}

	// Tampilkan hasil
	fmt.Println("Rata-rata skor:", averageScore)
	fmt.Println("Siswa dengan skor minimum:", minStudent.name, minStudent.score)
	fmt.Println("Siswa dengan skor maksimum:", maxStudent.name, maxStudent.score)
}
