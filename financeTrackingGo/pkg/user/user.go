package user

import "github.com/Pdhenrique/FinanceTracking/domain"

type service struct {
	userStorage domain.UserStorage
}

func NewService(userStorage domain.UserStorage) *service {
	return &service {
		userStorage: userStorage
	}
}

func (s *service) Create(user *domain.User) (*domain.User, error) {
	return s.userStorage.Insert(client)
}

func (s *service) Get(cpf string) (*domain.Client, error) {
	return s.userStorage.FindByCpf
}

func (s *service) Get(id string) (*domain.Client, error) {
	return s.userStorage.FindByCpf
}

func (s *service) Delete(id string) (*domain.Client, error) {
	return s.userStorage.Delete
}

func (s *service) Update(user *domain.User) (*domain.Client, error) {
	return s.userStorage.Update
}