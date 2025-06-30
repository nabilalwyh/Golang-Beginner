// Golang Beginner Pert 7 - File main.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":8187" // ganti 3 digit belakang port pake 3 digit belakang npm kalian
	// misal, npm: 51422189, berarti port nya 8187

	http.Handle("/", http.FileServer(http.Dir("polymer")))

	fmt.Println("Listening on http://localhost" + PORT)

	log.Fatal(http.ListenAndServe(PORT, nil))
}
