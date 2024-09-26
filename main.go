package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "PAYMENT_API: ", log.LstdFlags)

	// Crear un nuevo router con Gorilla Mux
	router := mux.NewRouter()


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
