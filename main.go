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

	// Crear un nuevo controlador
	controllerUser := controllers.NewUserController(logger)
	controllerRoom := controllers.NewRoomController(logger)
	controllerReservation := controllers.NewReservationController(logger)
	controllerPayment := controllers.NewPaymentController(logger)

	// Crear un nuevo router con Gorilla Mux
	router := mux.NewRouter()

	// Rutas de usuario
	router.HandleFunc("/user", controllerUser.Get).Methods("GET")
	router.HandleFunc("/user/{id}", controllerUser.GetID).Methods("GET")
	router.HandleFunc("/user", controllerUser.Post).Methods("POST")
	router.HandleFunc("/user/{id}", controllerUser.Modify).Methods("PUT")
	router.HandleFunc("/user/{id}", controllerUser.Delete).Methods("DELETE")

	// Rutas de habitacion
	router.HandleFunc("/habitacion", controllerRoom.Get).Methods("GET")
	router.HandleFunc("/habitacion/{id}", controllerRoom.GetID).Methods("GET")
	router.HandleFunc("/disponible", controllerRoom.GetAvailable).Methods("GET")
	router.HandleFunc("/habitacion", controllerRoom.Post).Methods("POST")
	router.HandleFunc("/habitacion/{id}", controllerRoom.Modify).Methods("PUT")
	router.HandleFunc("/habitacion/{id}", controllerRoom.Delete).Methods("DELETE")

	// Rutas de reserva
	router.HandleFunc("/reserva", controllerReservation.Get).Methods("GET")
	router.HandleFunc("/reserva/{id}", controllerReservation.GetID).Methods("GET")
	router.HandleFunc("/reserva", controllerReservation.Create).Methods("POST")
	router.HandleFunc("/reserva/{id}", controllerReservation.Mod).Methods("PUT")
	router.HandleFunc("/reserva/{id}", controllerReservation.Del).Methods("DELETE")

	// Rutas de pago
	router.HandleFunc("/pago", controllerPayment.Get).Methods("GET")
	router.HandleFunc("/pago/{id}", controllerPayment.GetID).Methods("GET")
	router.HandleFunc("/pago", controllerPayment.Create).Methods("POST")
	router.HandleFunc("/pago/{id}", controllerPayment.Mod).Methods("PUT")
	router.HandleFunc("/pago/{id}", controllerPayment.Del).Methods("DELETE")

	// Iniciar el servidor
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Iniciar el servidor en un goroutine
	go func() {
		logger.Println("Starting server on port :8080")
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
