package enum

type OrderStatus = uint8

const (
	WaitingPaymentApproval OrderStatus = iota
	PaymentApproved
	Delivered
	Finished
)

func IsValid(value int) bool {
	return value <= int(Finished)
}
