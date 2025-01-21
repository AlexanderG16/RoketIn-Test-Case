package handler

import (
	"challenge-2/backend/models"
	"challenge-2/backend/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateMovies(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var movie models.Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = services.CreateMovie(db, &movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(movie)
	}
}

func GetAllMovies(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		pages := r.URL.Query()["page"]
		page, err := strconv.Atoi(strings.Join(pages[:], ""))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		movies, err := services.GetAllMovies(db, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movies)
	}
}

func SearchMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		title := r.URL.Query()["title"]
		description := r.URL.Query()["description"]
		artists := r.URL.Query()["artists"]
		genres := r.URL.Query()["genres"]
		movies, err := services.SearchMovies(db, title, description, artists, genres)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movies)
	}
}

func UpdateMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var movie models.Movie
		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := services.UpdateMovie(db, &movie, uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movie)
	}
}
