package main

import "fmt"

func main() {
	// Bagian 1: Perhitungan nilai dan grade
	var uts, uas, total float32
	var grade string

	fmt.Print("Masukkan nilai UTS : ")
	fmt.Scanln(&uts)
	fmt.Print("Masukkan nilai UAS : ")
	fmt.Scanln(&uas)

	total = uts*0.3 + uas*0.7

	if total >= 81 && total <= 100 {
		grade = "A"
	} else if total >= 61 && total < 81 {
		grade = "B"
	} else if total >= 41 && total < 61 {
		grade = "C"
	} else if total >= 21 && total < 41 {
		grade = "D"
	} else if total <= 20 {
		grade = "E"
	} else {
		fmt.Println("Inputan salah!")
	}

	fmt.Println("\n======= Hasil nilai =======")
	fmt.Printf("Nilai UTS\t: %.2f \n", uts)
	fmt.Printf("Nilai UAS\t: %.2f\n", uas)
	fmt.Printf("Nilai\t\t: %.2f\n", total)
	fmt.Printf("Grade\t\t: %s\n", grade)
}
