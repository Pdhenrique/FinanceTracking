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

	err := u.DB.QueryRow(`SELECT * FROM tb_users WHERE id = $1`, id).Scan(
		&user.ID,
		&user.CPF,
		&user.NAME,
		&user.EMAIL,
		&user.PASSWORD,
		&user.CREATED_AT,
		&user.UPDATED_AT,
		&user.ACTIVE)

	if err != nil {
		log.Printf("Error finding user by ID %s: %v", id, err)
	}

	return &user, nil
}

func (u *userStorage) FindByCpf(cpf string) (*domain.User, error) {
	var user domain.User

	err := u.DB.QueryRow(`SELECT * FROM tb_users WHERE cpf = $1`, cpf).Scan(
		&user.ID,
		&user.CPF,
		&user.NAME,
		&user.EMAIL,
		&user.PASSWORD,
		&user.CREATED_AT,
		&user.UPDATED_AT,
		&user.ACTIVE)

	if err != nil {
		log.Printf("Error finding user by CPF %s: %v", cpf, err)
	}

	return &user, nil

}

func (u *userStorage) Delete(id string) error {
	_, err := u.DB.Exec(`DELETE FROM tb_users WHERE id = $1`, id)

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
		user.ACTIVE)

	if err != nil {
		log.Printf("Error updating user with ID %s: %v", user.ID, err)
	}

	return err
}

func (u *userStorage) Insert(user *domain.User) (*domain.User, error) {
	err := u.DB.QueryRow(`INSERT INTO tb_users (id, cpf, name, email, password, created_at, updated_at, ACTIVE) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		user.ID,
		user.CPF,
		user.NAME,
		user.EMAIL,
		user.PASSWORD,
		user.CREATED_AT,
		user.UPDATED_AT,
		user.ACTIVE).Scan(&user.ID)

	if err != nil {
		log.Printf("Error inserting user with ID %s: %v", user.ID, err)
		return nil, err
	}

	return user, nil
}

func (u *userStorage) FindByEmailOrCpf(email string, cpf string) (*domain.User, error) {
	var user domain.User

	err := u.DB.QueryRow(`SELECT * FROM tb_users WHERE email = $1 OR cpf = $2`, email, cpf).Scan(
		&user.ID,
		&user.CPF,
		&user.NAME,
		&user.EMAIL,
		&user.PASSWORD,
		&user.CREATED_AT,
		&user.UPDATED_AT,
		&user.ACTIVE)

	if err != nil {
		log.Printf("Error finding user by email %s or CPF %s: %v", email, cpf, err)
	}

	return &user, nil
}
