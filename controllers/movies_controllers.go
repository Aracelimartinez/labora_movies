package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Aracelimartinez/labora_movies/models"
	"github.com/Aracelimartinez/labora_movies/services"
)

// Función para crear películas
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie models.Movie

	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMovieName, err := services.CreateMovie(newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al crear la película"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := fmt.Sprintf("Película '%s' adicionada con éxito!", newMovieName)

	w.Write([]byte(response))
}
