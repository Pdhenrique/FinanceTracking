package account

import "github.com/Pdhenrique/FinanceTracking/domain"

type service struct {
	accountStorage domain.AccountStorage
}

func NewService(accountStorage domain.AccountStorage) *service {
	return &service{
		accountStorage: accountStorage,
	}
}

func (s *service) Delete(id string) error {
	return s.accountStorage.Delete(id)
}

func (s *service) Post(account *domain.Account) (*domain.Account, error) {
	return s.accountStorage.Insert(account)
}

func (s *service) Put(account *domain.Account) error {
	return s.accountStorage.Update(account)
}

func (s *service) GetByID(id string) (*domain.Account, error) {
	return s.accountStorage.FindByID(id)
}
