package db

import (
	"database/sql"
	"log"

	"github.com/Pdhenrique/FinanceTracking/domain"
)

type userStorage struct {
	DB *sql.DB
}

func NewUserStorage(db *sql.DB) domain.UserStorage {
	return &userStorage{
		DB: db,
	}
}

func (u *userStorage) FindByID(id string) (*domain.User, error) {
	var user domain.User

	err := u.DB.QueryRow(``, id).Scan(
		&user.ID,
		&user.CPF,
		&user.NAME,
		&user.EMAIL,
		&user.PASSWORD,
		&user.CREATED_AT,
		&user.UPDATED_AT,
		&user.ACITVE)

	if err != nil {
		log.Printf("Error finding user by ID %s: %v", id, err)
	}

	return &user, nil
}

func (u *userStorage) FindByCpf(cpf string) (*domain.User, error) {
	var user domain.User

	err := u.DB.QueryRow(``, cpf).Scan(
		&user.ID,
		&user.CPF,
		&user.NAME,
		&user.EMAIL,
		&user.PASSWORD,
		&user.CREATED_AT,
		&user.UPDATED_AT,
		&user.ACITVE)

	if err != nil {
		log.Printf("Error finding user by CPF %s: %v", cpf, err)
	}

	return &user, nil

}

func (u *userStorage) Delete(id string) error {
	_, err := u.DB.Exec(``, id)

	if err != nil {
		log.Printf("Error deleting user with ID %s: %v", id, err)
	}

	return err
}

func (u *userStorage) Update(user *domain.User) error {
	_, err := u.DB.Exec(``,
		user.ID,
		user.CPF,
		user.NAME,
		user.EMAIL,
		user.PASSWORD,
		user.CREATED_AT,
		user.UPDATED_AT,
		user.ACITVE)

	if err != nil {
		log.Printf("Error updating user with ID %s: %v", user.ID, err)
	}

	return err
}

func (u *userStorage) Insert(user *domain.User) (*domain.User, error) {
	err := u.DB.QueryRow(``,
		user.ID,
		user.CPF,
		user.NAME,
		user.EMAIL,
		user.PASSWORD,
		user.CREATED_AT,
		user.UPDATED_AT,
		user.ACITVE).Scan(&user.ID)

	if err != nil {
		log.Printf("Error inserting user with ID %s: %v", user.ID, err)
		return nil, err
	}

	return user, nil
}
