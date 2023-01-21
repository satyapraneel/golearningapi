package controllers

import (
	"encoding/json"
	"myapp/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := models.GetAllMovies()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	var movie models.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		json.NewEncoder(w).Encode("No data found?")
		return
	}

	models.CreateMovie(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	models.UpdateMovie(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	models.DeleteMovie(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	models.DeleteAllMovies()
}
