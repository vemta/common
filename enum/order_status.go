package enum

type OrderStatus = uint8

const (
	WaitingPaymentApproval OrderStatus = iota
	PaymentApproved
	Delivered
	Finished
)
