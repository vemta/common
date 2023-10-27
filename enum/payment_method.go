package enum

type PaymentMethod struct {
	ID                 int
	Name               string
	SupportedCountries string
}

var (
	PIX PaymentMethod = PaymentMethod{
		ID:                 0x01,
		Name:               "PIX",
		SupportedCountries: "BR",
	}

	Boleto PaymentMethod = PaymentMethod{
		ID:                 0x02,
		Name:               "Boleto",
		SupportedCountries: "BR",
	}

	PayPal PaymentMethod = PaymentMethod{
		ID:   0x03,
		Name: "PayPal",
	}

	TED PaymentMethod = PaymentMethod{
		ID:   0x02,
		Name: "TED",
	}

	AvailablePaymentMethods = [...]PaymentMethod{PIX, Boleto, PayPal, TED}
)

func GetPaymentMethod(id int) *PaymentMethod {
	for _, p := range AvailablePaymentMethods {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
