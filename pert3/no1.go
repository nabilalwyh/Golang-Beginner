package main

import "fmt"

func main() {
	// Bagian 1: Nested loop dari perulanganA.go
	fmt.Println("=== Pola Nested Loop ===")
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}

	// Bagian 2: Menampilkan angka ganjil dari 0–14
	fmt.Println("\n=== Angka Ganjil (0–14) ===")
	for i := 0; i < 15; i++ {
		if i%2 != 0 {
			fmt.Println("Angka", i)
		}
	}

	// Bagian 3: Menentukan angka ganjil/genap dari 1–10
	fmt.Println("\n=== Ganjil atau Genap (1–10) ===")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i, "adalah angka ganjil")
		} else {
			fmt.Println(i, "adalah angka genap")
		}
	}
}
