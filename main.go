package main

import (
	"log"
	"net/http"
	"os"

	controllers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Controllers"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "PAYMENT_API: ", log.LstdFlags)

	// Crear un nuevo router con Gorilla Mux
	router := mux.NewRouter()

	// Registrar la ruta para procesar pagos
	router.HandleFunc("/api/pay", controllers.PaymentHandler(logger)).Methods("POST")

	// Iniciar el servidor
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
