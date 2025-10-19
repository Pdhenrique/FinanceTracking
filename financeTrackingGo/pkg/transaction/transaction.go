package transaction

import (
	"fmt"
	"io"

	"github.com/Pdhenrique/FinanceTracking/domain"
	"github.com/Pdhenrique/FinanceTracking/internal/parser"
)

type service struct {
	transactionStorage domain.TransactionStorage
}

func NewService(transactionStorage domain.TransactionStorage) *service {
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

func (s *service) GetByID(id string) (*domain.Transaction, error) {
	return s.transactionStorage.FindByID(id)
}

func (s *service) GetAll() ([]*domain.Transaction, error) {
	return s.transactionStorage.FindAll()
}

func (s *service) ImportTransactions(file io.Reader) (int, error) {
	transactions, err := parser.Parse(file)
	if err != nil {
		return 0, fmt.Errorf("error parsing transactions: %w", err)
	}

	if err := s.transactionStorage.ImportTransactions(transactions); err != nil {
		return 0, fmt.Errorf("error importing transactions: %w", err)
	}

	return len(transactions), nil
}
