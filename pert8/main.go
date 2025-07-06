package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":8187" // ganti 3 digit belakang port pake 3 digit belakang npm kalian
	// misal, npm: 51422189, berarti port nya 8187

	http.Handle("/", http.FileServer(http.Dir("polymer")))
	http.HandleFunc("/api/mahasiswa", user)

	fmt.Println("Listening on http://localhost" + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

type lepkom struct {
	Nama   string `json:"nama_mahasiswa"`
	Kursus string `json:"kursus_mahasiswa"`
	Foto   string `json:"foto_mahasiswa"`
}

var data_mahasiswa = []lepkom{
	{
		Nama:   "Nabila",
		Kursus: "Oracle Inter",
		Foto:   "img/gambar1.jpg",
	},
	{
		Nama:   "Bilal",
		Kursus: "Funda Web",
		Foto:   "img/gambar1.jpg",
	},
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == http.MethodGet {
		result, err := json.Marshal(data_mahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}
}
