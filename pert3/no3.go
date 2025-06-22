package main

import "fmt"

func main() {
	var i int32
	for j := 1; j <= 10; j++ {
		fmt.Print("\nInput Nilai I = ")
		fmt.Scanln(&i)

		if i <= 10 {
			if i%2 == 0 {
				fmt.Println(i, " bilangan genap")
			} else {
				fmt.Println(i, " bilangan ganjil")
			}
		} else {
			fmt.Println("Pertanyaan selesai, karena angka I yang diinput sudah melebihi dari 10. Terimakasih")
			break
		}
	}
}
