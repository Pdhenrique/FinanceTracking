package user

import (
	"fmt"
	"time"

	"github.com/Pdhenrique/FinanceTracking/domain"
)

type service struct {
	userStorage domain.UserStorage
}

func NewService(userStorage domain.UserStorage) *service {
	return &service{
		userStorage: userStorage,
	}
}

func (s *service) Create(user *domain.User) (*domain.User, error) {

	existing, err := s.userStorage.FindByEmailOrCpf(user.EMAIL, user.CPF)
	if err != nil {
		return nil, fmt.Errorf("error checking existing user: %w", err)
	}

	if existing != nil {
		return nil, fmt.Errorf("user with email %s or cpf %s already exists", existing.EMAIL, existing.CPF)
	}

	user.CREATED_AT = time.Now().Format(time.RFC3339)

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
