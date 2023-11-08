package enum

type OrderStatus uint8

const (
	WaitingPaymentApproval OrderStatus = iota
	WaitingConfirmation
	ToBeShipped
	Shipped
	WaitingDeliverConfirmation
	Canceled
	ToReturn
	Completed
)

func IsOrderStatusValid(value int) bool {
	return value <= int(Completed)
}
