package controller

import (
	"api-book-go-postgres/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id, omiempty"`
	Message string `json:"message, omiempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Buku `json:"data"`
}

func TambahBuku(w http.ResponseWriter, r *http.Request) {
	var buku models.Buku

	err := json.NewDecoder(r.Body).Decode(&buku)
	if err != nil {
		log.Fatalf("Can't decode from requst body. %v", err)
	}

	insertID := models.TambahBuku(buku)

	res := response{
		ID:      insertID,
		Message: "Books data added!",
	}

	json.NewEncoder(w).Encode(res)
}

func AmbilBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Can't convert string to int. %v", err)
	}

	buku, err := models.AmbilSatuBuku(int64(id))
	if err != nil {
		log.Fatalf("Can't get book's data. %v", err)
	}

	json.NewEncoder(w).Encode(buku)
}

func AmbilSemuaBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	bukus, err := models.AmbilSemuaBuku()
	if err != nil {
		log.Fatalf("Can't get the data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = bukus

	json.NewEncoder(w).Encode(response)
}

func UpdateBuku(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Can't convert string to int. %v", err)
	}

	var buku models.Buku

	err = json.NewDecoder(r.Body).Decode(&buku)
	if err != nil {
		log.Fatalf("Can't decode the request body. %v", err)
	}

	updatedRows := models.UpdateBuku(int64(id), buku)

	msg := fmt.Sprintf("Successfuly updating the book. The updated book are %v rows/record", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func HapusBuku(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Can't convert string to int. %v", err)
	}

	deletedRows := models.HapusBuku(int64(id))

	msg := fmt.Sprintf("Successfuly deleting the book. Deleted books are %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
