package user

import "github.com/Pdhenrique/FinanceTracking/domain"

type service struct {
	userStorage domain.UserStorage
}

func NewService(userStorage domain.UserStorage) *service {
	return &service{
		userStorage: userStorage,
	}
}

func (s *service) Create(user *domain.User) (*domain.User, error) {
	return s.userStorage.Insert(user)
}

func (s *service) GetByCpf(cpf string) (*domain.User, error) {
	return s.userStorage.FindByCpf(cpf)
}

func (s *service) GetByID(id string) (*domain.User, error) {
	return s.userStorage.FindByID(id)
}

func (s *service) Delete(id string) error {
	return s.userStorage.Delete(id)
}

func (s *service) Update(user *domain.User) error {
	return s.userStorage.Update(user)
}