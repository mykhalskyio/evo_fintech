package entity

type Filter struct {
	TransactionId    int
	TerminalId       string
	Status           string
	PaymentType      string
	DatePost         []string
	PaymentNarrative string
}
