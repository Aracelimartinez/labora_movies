package services

import (
	"errors"

	"github.com/Aracelimartinez/labora_movies/models"
)

func CreateMovie(newMovie models.Movie) (string, error) {

	// Validación de los campos requeridos
	if newMovie.Name == "" || newMovie.Year == 0 || newMovie.Gender == "" || newMovie.Price == 0 {
		return "", errors.New("todos los campos  excepto 'adquired' y 'stock' son obligatórios")
	}

	// Preparar la sentencia SQL para insertar la película
	stmt, err := Db.Prepare("INSERT INTO movies (name, year, gender, adquired, stock, price) VALUES ($1, $2, $3, $4, $5, $6) RETURNING name")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	// Ejecutar la sentencia y obtiene el nombre de la nueva película
	var newMovieName string
	err = stmt.QueryRow(newMovie.Name, newMovie.Year, newMovie.Gender, newMovie.Adquired, newMovie.Stock, newMovie.Price).Scan(&newMovieName)
	if err != nil {
		return "", err
	}

	return newMovieName, nil
}
