package payment

type Paypal struct {
	Email string
}

func (p *Paypal) Pay(amount float64) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": "Payment made with PayPal",
		"amount":  amount,
		"email":   p.Email,
	}
}
