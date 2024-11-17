package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	database "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Database"
	Middlewares_cors "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Middlewares"
	routes "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Routes"
	"github.com/gorilla/mux"
)

func main() {

	// Iniciar la base de datos
	database.Init()

	// Crear un nuevo logger
	logger := log.New(os.Stdout, "HOTEL-API: ", log.LstdFlags)

	// Crear un nuevo router con Gorilla Mux
	router := mux.NewRouter()

	// Rutas
	routes.AdminRoutes(router, logger)
	routes.UserRoutes(router, logger)
	routes.CompanionRoutes(router, logger)
	routes.RoomRoutes(router, logger)
	routes.ReservationRoutes(router, logger)
	routes.PaymentRoutes(router, logger)

	// Rutas de middlewares
	cors := Middlewares_cors.CorsMiddleware(router)

	// Iniciar el servidor
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      cors,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Iniciar el servidor en un goroutine
	go func() {
		logger.Println("Starting server in http://localhost:8080")
		err := srv.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Iniciar la funcion de shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	logger.Println("Received terminated, graceful shutdown", sig)
	tc, err := context.WithTimeout(context.Background(), 30*time.Second)

	// Cerrar el servidor
	if err != nil {
		logger.Println(err)
	}
	srv.Shutdown(tc)
}
