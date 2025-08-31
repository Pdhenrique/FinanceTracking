package domain

type Statement struct {
	BANK_NAME    string        `json:"bank_name"`
	AGENCY       string        `json:"agency"`
	ACCOUNT_ID   string        `json:"account_id"`
	GENERATED_AT string        `json:"generated_at"`
	PERIOD_START string        `json:"period_start"`
	PERIOD_END   string        `json:"period_end"`
	TRANSACTIONS []Transaction `json:"transactions"`
}

type StatementService interface {
	ImportStatement(statement []*Statement) error
	getByAcc(id string) (*Statement, error)
}

type StatementStorage interface {
	ImportStatement(Statement []*Statement) error
	findByAcc(id string) (*Statement, error)
}

func NewStatement(
	bankName string,
	agency string,
	accountId string,
	generatedAt string,
	periodStart string,
	periodEnd string,
	transactions []Transaction,
) *Statement {
	return &Statement{
		BANK_NAME:    bankName,
		AGENCY:       agency,
		ACCOUNT_ID:   accountId,
		GENERATED_AT: generatedAt,
		PERIOD_START: periodStart,
		PERIOD_END:   periodEnd,
		TRANSACTIONS: transactions,
	}
}
