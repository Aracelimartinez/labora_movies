package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Aracelimartinez/labora_movies/controllers"
	"github.com/Aracelimartinez/labora_movies/services"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	err := services.EstablishDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	//Definición de rutas
	router.HandleFunc("/", func(respuesta http.ResponseWriter, peticion *http.Request) {
		respuesta.Write([]byte("Hola Mundo!"))
	}).Methods(http.MethodGet)

	router.HandleFunc("/movies", controllers.CreateMovie).Methods(http.MethodPost)

	//Configuración de CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5432"},
		AllowedMethods: []string{"GET", "POST"},
	})

	handler := corsOptions.Handler(router)
	port := ":8000"

	if err := startServer(port, handler); err != nil {
		log.Fatalf("Error al iniciar el servidor: \n%v", err)
	}

}

func startServer(port string, router http.Handler) error {
	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Iniciando servidor en el puerto %s ..\n", port)
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
