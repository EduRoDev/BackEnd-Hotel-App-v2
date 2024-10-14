package routes

import (
	"log"
	"net/http"

	controllers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Controllers"
	"github.com/gorilla/mux"
)

func wrapHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		h(w, r)
	}
}

func UserRoutes(router *mux.Router, logger *log.Logger) {
	userController := controllers.NewUserController(logger)
	router.HandleFunc("/login", wrapHandler(userController.Login)).Methods("POST")
	router.HandleFunc("/user", wrapHandler(userController.Get)).Methods("GET")
	router.HandleFunc("/user/{id}", wrapHandler(userController.GetID)).Methods("GET")
	router.HandleFunc("/users/last", wrapHandler(userController.GetLastUser)).Methods("GET")
	router.HandleFunc("/user", wrapHandler(userController.Post)).Methods("POST")
	router.HandleFunc("/user/{id}", wrapHandler(userController.Modify)).Methods("PUT")
	router.HandleFunc("/user/{id}", wrapHandler(userController.Delete)).Methods("DELETE")
}

func RoomRoutes(router *mux.Router, logger *log.Logger) {
	roomController := controllers.NewRoomController(logger)
	router.HandleFunc("/habitacion", wrapHandler(roomController.Get)).Methods("GET")
	router.HandleFunc("/habitacion/{id}", wrapHandler(roomController.GetID)).Methods("GET")
	router.HandleFunc("/disponible", wrapHandler(roomController.GetAvailable)).Methods("GET")
	router.HandleFunc("/habitacion", wrapHandler(roomController.Post)).Methods("POST")
	router.HandleFunc("/habitacion/{id}", wrapHandler(roomController.Modify)).Methods("PUT")
	router.HandleFunc("/habitacion/{id}", wrapHandler(roomController.Delete)).Methods("DELETE")
}

func ReservationRoutes(router *mux.Router, logger *log.Logger) {
	reservationController := controllers.NewReservationController(logger)
	router.HandleFunc("/reserva", wrapHandler(reservationController.Get)).Methods("GET")
	router.HandleFunc("/reserva/{id}", wrapHandler(reservationController.GetID)).Methods("GET")
	router.HandleFunc("/reserva", wrapHandler(reservationController.Create)).Methods("POST")
	router.HandleFunc("/reserva/{id}", wrapHandler(reservationController.Mod)).Methods("PUT")
	router.HandleFunc("/cancelarReserva/{id}", wrapHandler(reservationController.Cancel)).Methods("DELETE")
	router.HandleFunc("/reserva/{id}", wrapHandler(reservationController.Del)).Methods("DELETE")
}

func PaymentRoutes(router *mux.Router, logger *log.Logger) {
	paymentController := controllers.NewPaymentController(logger)
	router.HandleFunc("/pago", wrapHandler(paymentController.Get)).Methods("GET")
	router.HandleFunc("/pago/{id}", wrapHandler(paymentController.GetID)).Methods("GET")
	router.HandleFunc("/pago", wrapHandler(paymentController.Create)).Methods("POST")
	router.HandleFunc("/pago/{id}", wrapHandler(paymentController.Mod)).Methods("PUT")
	router.HandleFunc("/pago/{id}", wrapHandler(paymentController.Del)).Methods("DELETE")
}

func KeyRoutes(router *mux.Router, logger *log.Logger) {
	keyController := controllers.NewKeyController(logger)
	router.HandleFunc("/llave", wrapHandler(keyController.Get)).Methods("GET")
	router.HandleFunc("/llave/{id}", wrapHandler(keyController.GetID)).Methods("GET")
	router.HandleFunc("/llave", wrapHandler(keyController.Create)).Methods("POST")
	router.HandleFunc("/llave/{id}", wrapHandler(keyController.Modify)).Methods("PUT")
	router.HandleFunc("/llave/{id}", wrapHandler(keyController.Delete)).Methods("DELETE")
}

func PersonalRoutes(router *mux.Router, logger *log.Logger) {
	personalController := controllers.NewPersonalController(logger)
	router.HandleFunc("/personal", wrapHandler(personalController.Get)).Methods("GET")
	router.HandleFunc("/personal/{id}", wrapHandler(personalController.GetID)).Methods("GET")
	router.HandleFunc("/personal", wrapHandler(personalController.Asign)).Methods("POST")
	router.HandleFunc("/personal/{id}", wrapHandler(personalController.Modify)).Methods("PUT")
	router.HandleFunc("/personal/{id}", wrapHandler(personalController.Delete)).Methods("DELETE")
}

func PersonalRoomRoutes(router *mux.Router, logger *log.Logger) {
	personalRoomController := controllers.NewPersonalRoomController(logger)
	router.HandleFunc("/personalRoom", wrapHandler(personalRoomController.Get)).Methods("GET")
	router.HandleFunc("/personalRoom/{id}", wrapHandler(personalRoomController.GetID)).Methods("GET")
	router.HandleFunc("/personalRoom", wrapHandler(personalRoomController.Asign)).Methods("POST")
	router.HandleFunc("/personalRoom/{id}", wrapHandler(personalRoomController.Modify)).Methods("PUT")
	router.HandleFunc("/personalRoom/{id}", wrapHandler(personalRoomController.Delete)).Methods("DELETE")
}

func InvoiceRoutes(router *mux.Router, logger *log.Logger) {
	invoiceController := controllers.NewInvoiceController(logger)
	router.HandleFunc("/invoice", wrapHandler(invoiceController.Get)).Methods("GET")
	router.HandleFunc("/invoice/{id}", wrapHandler(invoiceController.GetID)).Methods("GET")
	router.HandleFunc("/invoice", wrapHandler(invoiceController.Asign)).Methods("POST")
	router.HandleFunc("/invoice/{id}", wrapHandler(invoiceController.Modify)).Methods("PUT")
	router.HandleFunc("/invoice/{id}", wrapHandler(invoiceController.Delete)).Methods("DELETE")
}

func CheckInRoutes(router *mux.Router, logger *log.Logger) {
	checkInController := controllers.NewCheckController(logger)
	router.HandleFunc("/checkIn", wrapHandler(checkInController.Get)).Methods("GET")
	router.HandleFunc("/checkIn/{id}", wrapHandler(checkInController.GetID)).Methods("GET")
	router.HandleFunc("/checkIn", wrapHandler(checkInController.Asign)).Methods("POST")
	router.HandleFunc("/checkIn/{id}", wrapHandler(checkInController.Modify)).Methods("PUT")
	router.HandleFunc("/checkIn/{id}", wrapHandler(checkInController.Delete)).Methods("DELETE")
}
