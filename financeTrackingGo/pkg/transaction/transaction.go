package transaction

import "github.com/Pdhenrique/FinanceTracking/domain"

type service struct {
	transactionStorage domain.TransactionStorage
}

func NewService(transactionStorage domain.Transaction) *service {
	return &service{
		transactionStorage: transactionStorage,
	}
}

func (s *service) Delete(id string) error {
	return s.transactionStorage.Delete(id)
}

func (s *service) Post(transaction *domain.Transaction) (*domain.Transaction, error) {
	return s.transactionStorage.Insert(transaction)
}

func (s *service) Put(transaction *domain.Transaction) error {
	return s.transactionStorage.Update(transaction)
}

func (s *service) findByID(id string) (*domain.Transaction, error) {
	return s.transactionStorage.FindByID(id)
}

func (s *service) ImportTransactions(transactions []*domain.Transaction) error {
	return s.transactionStorage.ImportTransactions(transactions)
}
