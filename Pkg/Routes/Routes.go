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
	router.HandleFunc("/user/{id:[0-9]+}", wrapHandler(userController.GetID)).Methods("GET")
	router.HandleFunc("/user/{nombre}/{apellido}", wrapHandler(userController.GetUser)).Methods("GET")
	router.HandleFunc("/user", wrapHandler(userController.Post)).Methods("POST")
	router.HandleFunc("/user/{id:[0-9]+}", wrapHandler(userController.Modify)).Methods("PUT")
	router.HandleFunc("/user/{id:[0-9]+}", wrapHandler(userController.Delete)).Methods("DELETE")
}

func CompanionRoutes(router *mux.Router, logger *log.Logger) {
	companionController := controllers.NewCompanionController(logger)
	router.HandleFunc("/acompañante", wrapHandler(companionController.Get)).Methods("GET")
	router.HandleFunc("/acompañante/{id:[0-9]+}", wrapHandler(companionController.GetID)).Methods("GET")
	router.HandleFunc("/acompañante", wrapHandler(companionController.POST)).Methods("POST")
	router.HandleFunc("/acompañante/{id:[0-9]+}", wrapHandler(companionController.Mod)).Methods("PUT")
	router.HandleFunc("/acompañante/{id:[0-9]+}", wrapHandler(companionController.Delete)).Methods("DELETE")
}

func RoomRoutes(router *mux.Router, logger *log.Logger) {
	roomController := controllers.NewRoomController(logger)
	router.HandleFunc("/habitacion", wrapHandler(roomController.Get)).Methods("GET")
	router.HandleFunc("/habitacion/{id:[0-9]+}", wrapHandler(roomController.GetID)).Methods("GET")
	router.HandleFunc("/disponible", wrapHandler(roomController.GetAvailable)).Methods("GET")
	router.HandleFunc("/habitacion", wrapHandler(roomController.Post)).Methods("POST")
	router.HandleFunc("/habitacion/{id:[0-9]+}", wrapHandler(roomController.Modify)).Methods("PUT")
	router.HandleFunc("/habitacion/{id:[0-9]+}", wrapHandler(roomController.Delete)).Methods("DELETE")
}

func ReservationRoutes(router *mux.Router, logger *log.Logger) {
	reservationController := controllers.NewReservationController(logger)
	router.HandleFunc("/reserva", wrapHandler(reservationController.Get)).Methods("GET")
	router.HandleFunc("/reserva/{id:[0-9]+}", wrapHandler(reservationController.GetID)).Methods("GET")
	router.HandleFunc("/reserva/{idUsuario}/{fechaEntrada}", wrapHandler(reservationController.GetByUsuarioYFecha)).Methods("GET")
	router.HandleFunc("/reserva", wrapHandler(reservationController.Create)).Methods("POST")
	router.HandleFunc("/reserva/{id:[0-9]+}", wrapHandler(reservationController.Mod)).Methods("PUT")
	router.HandleFunc("/reserva/{id:[0-9]+}", wrapHandler(reservationController.Del)).Methods("DELETE")
}

func PaymentRoutes(router *mux.Router, logger *log.Logger) {
	paymentController := controllers.NewPaymentController(logger)
	router.HandleFunc("/pago", wrapHandler(paymentController.Get)).Methods("GET")
	router.HandleFunc("/pago/{id:[0-9]+}", wrapHandler(paymentController.GetID)).Methods("GET")
	router.HandleFunc("/pago/reserva/{id:[0-9]+}", wrapHandler(paymentController.GetByIdReserva)).Methods("GET")
	router.HandleFunc("/pago", wrapHandler(paymentController.Create)).Methods("POST")
	router.HandleFunc("/pago/{id:[0-9]+}", wrapHandler(paymentController.Mod)).Methods("PUT")
	router.HandleFunc("/pago/{id:[0-9]+}", wrapHandler(paymentController.Del)).Methods("DELETE")
}

func PersonalRoutes(router *mux.Router, logger *log.Logger) {
	personalController := controllers.NewPersonalController(logger)
	router.HandleFunc("/personal", wrapHandler(personalController.Get)).Methods("GET")
	router.HandleFunc("/personal/{id:[0-9]+}", wrapHandler(personalController.GetID)).Methods("GET")
	router.HandleFunc("/personal", wrapHandler(personalController.Asign)).Methods("POST")
	router.HandleFunc("/personal/{id:[0-9]+}", wrapHandler(personalController.Modify)).Methods("PUT")
	router.HandleFunc("/personal/{id:[0-9]+}", wrapHandler(personalController.Delete)).Methods("DELETE")
}

func PersonalRoomRoutes(router *mux.Router, logger *log.Logger) {
	personalRoomController := controllers.NewPersonalRoomController(logger)
	router.HandleFunc("/personalRoom", wrapHandler(personalRoomController.Get)).Methods("GET")
	router.HandleFunc("/personalRoom/{id:[0-9]+}", wrapHandler(personalRoomController.GetID)).Methods("GET")
	router.HandleFunc("/personalRoom", wrapHandler(personalRoomController.Asign)).Methods("POST")
	router.HandleFunc("/personalRoom/{id:[0-9]+}", wrapHandler(personalRoomController.Modify)).Methods("PUT")
	router.HandleFunc("/personalRoom/{id:[0-9]+}", wrapHandler(personalRoomController.Delete)).Methods("DELETE")
}
