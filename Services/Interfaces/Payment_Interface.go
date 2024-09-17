package interfaces

type PaymentStrategy interface {
	Pay(amount float64) map[string]interface{}
}