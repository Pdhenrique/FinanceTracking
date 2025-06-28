package domain

type Transaction struct {
	ID string `json:"id"`
}

type TransactionService interface {
  Create(transaction *Transaction) (*Transaction, error)
  Update(transaction *Transaction) error
  Delete(id string) error
}

type TransactionStorage interface {
  Insert(transaction *Transaction) (*Transaction, error)
  Update(transaction *Transaction) error
  Delete(id string) error
}

func NewTransaction(
  id string,
) *User {
  return &User{
    ID: id
  }
}