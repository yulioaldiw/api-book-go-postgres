package router

import (
	"api-book-go-postgres/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/buku", controller.AmbilSemuaBuku).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controller.AmbilBuku).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/buku", controller.TambahBuku).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controller.UpdateBuku).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/buku/{id}", controller.HapusBuku).Methods("DELETE", "OPTIONS")

	return router
}
