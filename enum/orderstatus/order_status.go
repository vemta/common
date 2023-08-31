package orderstatus

type OrderStatus uint8

const (
	WaitingPaymentApproval OrderStatus = iota
	PaymentApproved
	Delivered
	Finished
)

func IsOrderStatusValid(value int) bool {
	return value <= int(Finished)
}
