package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	controllers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Controllers"
	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	"github.com/gorilla/mux"
)

func main() {
	// Iniciar la base de datos
	database.Init()

	// Crear un nuevo logger
	logger := log.New(os.Stdout, "HOTEL-API: ", log.LstdFlags)

	// Crear un nuevo controlador de usuario
	controller := controllers.NewUserController(logger)

	// Crear un nuevo router con Gorilla Mux
	router := mux.NewRouter()

	// Rutas de usuario
	router.HandleFunc("/user", controller.Get).Methods("GET")
	router.HandleFunc("/user/{id}", controller.GetID).Methods("GET")
	router.HandleFunc("/user", controller.Post).Methods("POST")
	router.HandleFunc("/user/{id}", controller.Modify).Methods("PUT")
	router.HandleFunc("/user/{id}", controller.Delete).Methods("DELETE")

	// Iniciar el servidor
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port :8080")
		err := srv.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	logger.Println("Received terminated, graceful shutdown", sig)
	tc, err := context.WithTimeout(context.Background(), 30*time.Second)

	if err != nil {
		logger.Println(err)
	}
	srv.Shutdown(tc)
}
