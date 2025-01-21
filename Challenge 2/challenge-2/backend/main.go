package main

import (
	"challenge-2/backend/config"
	"challenge-2/backend/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/movies", handler.GetAllMovies(config.DB)).Methods("GET")
	router.HandleFunc("/movies", handler.CreateMovies(config.DB)).Methods("POST")
	router.HandleFunc("/movies/{id}", handler.UpdateMovie(config.DB)).Methods("PUT")
	router.HandleFunc("/movies/search", handler.SearchMovie(config.DB)).Methods("GET")

	log.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", router)

}
