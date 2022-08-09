package main

import (
	"fmt"
	"log"
	"net/http"

	mahasiswacontroller "github.com/xvbnm48/go-crud-dasar/controllers/mahasiswa_controller"
)

func main() {
	http.HandleFunc("/", mahasiswacontroller.Index)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
