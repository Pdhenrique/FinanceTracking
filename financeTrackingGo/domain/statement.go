package domain

type Statement struct {
	BankName     string        `json:"bank_name"`
	Agency       string        `json:"agency"`
	AccountID    string        `json:"account_id"`
	GeneratedAt  string        `json:"generated_at"`
	PeriodStart  string        `json:"period_start"`
	PeriodEnd    string        `json:"period_end"`
	Transactions []Transaction `json:"transactions"`
}
