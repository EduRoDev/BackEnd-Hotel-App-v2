package routes

import (
	"log"
	"net/http"

	controllers "github.com/EduRoDev/BackEnd-Hotel-App-v2/Pkg/Controllers"
	"github.com/gorilla/mux"
)

var prefix = "/api/v1"

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
	router.HandleFunc(prefix+"/user/login", wrapHandler(userController.Login)).Methods("POST")
	router.HandleFunc(prefix+"/user", wrapHandler(userController.Get)).Methods("GET")
	router.HandleFunc(prefix+"/user/{id:[0-9]+}", wrapHandler(userController.GetID)).Methods("GET")
	router.HandleFunc(prefix+"/user/{nombre}/{apellido}", wrapHandler(userController.GetUser)).Methods("GET")
	router.HandleFunc(prefix+"/user", wrapHandler(userController.Post)).Methods("POST")
	router.HandleFunc(prefix+"/user/{id:[0-9]+}", wrapHandler(userController.Modify)).Methods("PUT")
	router.HandleFunc(prefix+"/user/{id:[0-9]+}", wrapHandler(userController.Delete)).Methods("DELETE")
}

func CompanionRoutes(router *mux.Router, logger *log.Logger) {
	companionController := controllers.NewCompanionController(logger)
	router.HandleFunc(prefix+"/acompañante", wrapHandler(companionController.Get)).Methods("GET")
	router.HandleFunc(prefix+"/acompañante/{id:[0-9]+}", wrapHandler(companionController.GetID)).Methods("GET")
	router.HandleFunc(prefix+"/acompañante", wrapHandler(companionController.POST)).Methods("POST")
	router.HandleFunc(prefix+"/acompañante/{id:[0-9]+}", wrapHandler(companionController.Mod)).Methods("PUT")
	router.HandleFunc(prefix+"/acompañante/{id:[0-9]+}", wrapHandler(companionController.Delete)).Methods("DELETE")
}

func RoomRoutes(router *mux.Router, logger *log.Logger) {
	roomController := controllers.NewRoomController(logger)
	router.HandleFunc(prefix+"/habitacion", wrapHandler(roomController.Get)).Methods("GET")
	router.HandleFunc(prefix+"/habitacion/{id:[0-9]+}", wrapHandler(roomController.GetID)).Methods("GET")
	router.HandleFunc(prefix+"/disponible", wrapHandler(roomController.GetAvailable)).Methods("GET")
	router.HandleFunc(prefix+"/habitacion", wrapHandler(roomController.Post)).Methods("POST")
	router.HandleFunc(prefix+"/habitacion/{id:[0-9]+}", wrapHandler(roomController.Modify)).Methods("PUT")
	router.HandleFunc(prefix+"/habitacion/{id:[0-9]+}", wrapHandler(roomController.Delete)).Methods("DELETE")
}

func ReservationRoutes(router *mux.Router, logger *log.Logger) {
	reservationController := controllers.NewReservationController(logger)
	router.HandleFunc(prefix+"/reserva", wrapHandler(reservationController.Get)).Methods("GET")
	router.HandleFunc(prefix+"/reserva/{id:[0-9]+}", wrapHandler(reservationController.GetID)).Methods("GET")
	router.HandleFunc(prefix+"/reserva/{idUsuario}/{fechaEntrada}", wrapHandler(reservationController.GetByUsuarioYFecha)).Methods("GET")
	router.HandleFunc(prefix+"/reserva", wrapHandler(reservationController.Create)).Methods("POST")
	router.HandleFunc(prefix+"/reserva/{id:[0-9]+}", wrapHandler(reservationController.Mod)).Methods("PUT")
	router.HandleFunc(prefix+"/reserva/{id:[0-9]+}", wrapHandler(reservationController.Del)).Methods("DELETE")
}

func PaymentRoutes(router *mux.Router, logger *log.Logger) {
	paymentController := controllers.NewPaymentController(logger)
	router.HandleFunc(prefix+"/pago", wrapHandler(paymentController.Get)).Methods("GET")
	router.HandleFunc(prefix+"/pago/{id:[0-9]+}", wrapHandler(paymentController.GetID)).Methods("GET")
	router.HandleFunc(prefix+"/pago/reserva/{id:[0-9]+}", wrapHandler(paymentController.GetByIdReserva)).Methods("GET")
	router.HandleFunc(prefix+"/pago", wrapHandler(paymentController.Create)).Methods("POST")
	router.HandleFunc(prefix+"/pago/{id:[0-9]+}", wrapHandler(paymentController.Mod)).Methods("PUT")
	router.HandleFunc(prefix+"/pago/{id:[0-9]+}", wrapHandler(paymentController.Del)).Methods("DELETE")
}

func AdminRoutes(router *mux.Router, logger *log.Logger) {
	adminController := controllers.NewAdminController(logger)
	router.HandleFunc(prefix+"/admin/login", wrapHandler(adminController.Login)).Methods("POST")
	router.HandleFunc(prefix+"/admin", wrapHandler(adminController.Create)).Methods("POST")
}
