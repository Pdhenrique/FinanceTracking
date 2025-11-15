package db

import (
	"database/sql"
	"log"

	"github.com/Pdhenrique/FinanceTracking/domain"
)

type accountStorage struct {
	DB *sql.DB
}

func NewAccountStorage(db *sql.DB) *accountStorage {
	return &accountStorage{
		DB: db,
	}
}

func (a *accountStorage) Insert(account *domain.Account) (*domain.Account, error) {
	err := a.DB.QueryRow(
		`INSERT INTO tb_accounts (id, account_type, active, balance, bank_name, created_at, updated_at, user_id) 
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
		&account.ID,
		&account.ACCOUNT_TYPE,
		&account.ACTIVE,
		&account.BALANCE,
		&account.BANK_NAME,
		&account.CREATED_AT,
		&account.UPDATED_AT,
		&account.USER_ID,
	).Scan(&account.ID)

	if err != nil {
		log.Printf("Error inserting account with ID %s: %v", account.ID, err)
		return nil, err
	}

	return account, nil
}

func (a *accountStorage) Update(account *domain.Account) error {
	_, err := a.DB.Exec(``,
		&account.ID,
		&account.ACCOUNT_TYPE,
		&account.ACTIVE,
		&account.BALANCE,
		&account.BANK_NAME,
		&account.CREATED_AT,
		&account.UPDATED_AT,
		&account.USER_ID,
	)

	if err != nil {
		log.Printf("Error updating account with ID %s: %v", account.ID, err)
	}

	return err
}

func (a *accountStorage) FindByID(id string) (*domain.Account, error) {
	account := &domain.Account{}

	err := a.DB.QueryRow(
		`SELECT id, account_type, active, balance, bank_name, created_at, updated_at, user_id
                 FROM tb_accounts WHERE id = $1`,
		id,
	).Scan(
		&account.ID,
		&account.ACCOUNT_TYPE,
		&account.ACTIVE,
		&account.BALANCE,
		&account.BANK_NAME,
		&account.CREATED_AT,
		&account.UPDATED_AT,
		&account.USER_ID,
	)

	if err != nil {
		log.Printf("Error finding account with ID %s: %v", id, err)
		return nil, err
	}

	return account, nil
}

func (a *accountStorage) Delete(id string) error {
	_, err := a.DB.Exec(``, id)

	if err != nil {
		log.Printf("Error deleting account with ID: %s: %v", id, err)
	}

	return err
}
