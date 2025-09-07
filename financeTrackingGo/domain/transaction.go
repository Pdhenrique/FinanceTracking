package domain

import "io"

type Transaction struct {
	ID              string  `json:"id"`
	AGENCY          string  `json:"agency"`
	ACCOUNT_ID      string  `json:"account_id"`
	RELEASE_DATE    string  `json:"release_date"`
	ACCOUNTING_DATE string  `json:"accounting_date"`
	TITLE           string  `json:"title"`
	DESCRIPTION     string  `json:"description"`
	INCOME          float64 `json:"income"`
	EXPENSE         float64 `json:"expense"`
	DAILY_BALANCE   float64 `json:"daily_balance"`
}

type TransactionService interface {
	Post(transaction *Transaction) (*Transaction, error)
	Put(transaction *Transaction) error
	Delete(id string) error
	ImportTransactions(r io.Reader) (int, error)
	GetByID(id string) (*Transaction, error)
}

type TransactionStorage interface {
	Insert(transaction *Transaction) (*Transaction, error)
	Update(transaction *Transaction) error
	Delete(id string) error
	ImportTransactions(transactions []*Transaction) error
	FindByID(id string) (*Transaction, error)
}

func NewTransaction(
	id string,
	agency string,
	accountID string,
	releaseDate string,
	accountingDate string,
	title string,
	description string,
	income float64,
	expense float64,
	dailyBalance float64,
) *Transaction {
	return &Transaction{
		ID:              id,
		AGENCY:          agency,
		ACCOUNT_ID:      accountID,
		RELEASE_DATE:    releaseDate,
		ACCOUNTING_DATE: accountingDate,
		TITLE:           title,
		DESCRIPTION:     description,
		INCOME:          income,
		EXPENSE:         expense,
		DAILY_BALANCE:   dailyBalance,
	}
}
