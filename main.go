package main

import (
	"api-book-go-postgres/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Server dijalankan pada port 8000...")

	log.Fatal(http.ListenAndServe(":8000", r))
}
