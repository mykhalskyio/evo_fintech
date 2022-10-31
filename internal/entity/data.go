package entity

type Data struct {
	TransactionId      int     `db:"transaction_id" json:"transaction_id"`
	RequestId          int     `db:"request_id" json:"request_id"`
	TerminalId         int     `db:"terminal_id" json:"terminal_id"`
	PartnerObjectId    int     `db:"partner_object_id" json:"partner_object_id"`
	AmountTotal        float32 `db:"amount_total" json:"amount_total"`
	AmountOriginal     float32 `db:"amount_original" json:"amount_original"`
	CommissionPS       float32 `db:"commission_ps" json:"commission_ps"`
	CommissionClient   float32 `db:"commission_client" json:"commission_client"`
	CommissionProvider float32 `db:"commission_provider" json:"commission_provider"`
	DateInput          string  `db:"date_input" json:"date_input"`
	DatePost           string  `db:"date_post" json:"date_post"`
	Status             string  `db:"status" json:"status"`
	PaymentType        string  `db:"payment_type" json:"payment_type"`
	PaymentNumber      string  `db:"payment_number" json:"payment_number"`
	ServiceId          int     `db:"service_id" json:"service_id"`
	Service            string  `db:"service" json:"service"`
	PayeeId            int     `db:"payee_id" json:"payee_id"`
	PayeeName          string  `db:"payee_name" json:"payee_name"`
	PayeeBankMfo       int     `db:"payee_bank_mfo" json:"payee_bank_mfo"`
	PayeeBankAccount   string  `db:"payee_bank_account" json:"payee_bank_account"`
	PaymentNarrative   string  `db:"payment_narrative" json:"payment_narrative"`
}
