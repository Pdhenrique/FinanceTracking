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
		DB: db
	}
}

func (u *userStorage) FindById(id string) (*domain.User, error){
	var client domain.client

	err := u.DB.QueryRow()
}