package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	payment "github.com/EduRoDev/BackEnd-Hotel-App-v2/Services/Impl/Payment"
	interfaces "github.com/EduRoDev/BackEnd-Hotel-App-v2/Services/Interfaces"
)

type PaymentRequest struct {
	PaymentType string  `json:"payment_type"`
	Amount      float64 `json:"amount"`
	Email       string  `json:"email"`
}

type PaymentLogger struct {
	l *log.Logger
	p interfaces.PaymentStrategy
}

func NewPaymentLogger(l *log.Logger, p interfaces.PaymentStrategy) *PaymentLogger {
	return &PaymentLogger{l, p}
}

func (p *PaymentLogger) PayProcess(paymentType string, amount float64) map[string]interface{} {
	var result map[string]interface{}
	switch paymentType {
	case "paypal":
		result = p.p.Pay(amount)
	default:
		result = map[string]interface{}{
			"status":  "error",
			"message": "Payment type not supported",
		}
	}

	p.l.Println(result)
	return result
}

func PaymentHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PaymentRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		var strategy interfaces.PaymentStrategy
		switch req.PaymentType {
		case "paypal":
			strategy = &payment.Paypal{Email: req.Email}
		default:
			http.Error(w, "Payment type not supported", http.StatusBadRequest)
			return
		}

		paymentLogger := NewPaymentLogger(logger, strategy)

		result := paymentLogger.PayProcess(req.PaymentType, req.Amount)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
