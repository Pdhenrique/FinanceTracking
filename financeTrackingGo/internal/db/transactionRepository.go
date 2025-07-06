package db

import (
	"database/sql"
	"log"

	"github.com/Pdhenrique/FinanceTracking/domain"
)

type transactionStorage struct {
	DB *sql.DB
}

func NewTransactionStorage(db *sql.DB) *transactionStorage {
	return &transactionStorage{
		DB: db,
	}
}

func (t *transactionStorage) Post(transaction *domain.Transaction) (*domain.Transaction, error) {
	err := t.DB.QueryRow(
		`INSERT INTO tb_transactions (id, agency, account_id, release_date, accounting_date, title, description, income, expense, daily_balance) 
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		&transaction.ID,
		&transaction.AGENCY,
		&transaction.ACCOUNT_ID,
		&transaction.RELEASE_DATE,
		&transaction.ACCOUNTING_DATE,
		&transaction.TITLE,
		&transaction.DESCRIPTION,
		&transaction.INCOME,
		&transaction.EXPENSE,
		&transaction.DAILY_BALANCE,
	).Scan(&transaction.ID)

	if err != nil {
		log.Printf("Error inserting transaction with ID %s: %v", transaction.ID, err)
		return nil, err
	}

	return transaction, nil
}

func (t *transactionStorage) Put(transaction *domain.Transaction) error {
	_, err := t.DB.Exec(``,
		&transaction.ID,
		&transaction.AGENCY,
		&transaction.ACCOUNT_ID,
		&transaction.RELEASE_DATE,
		&transaction.ACCOUNTING_DATE,
		&transaction.TITLE,
		&transaction.DESCRIPTION,
		&transaction.INCOME,
		&transaction.EXPENSE,
		&transaction.DAILY_BALANCE,
	)

	if err != nil {
		log.Printf("Error updating transaction with ID %s: %v", transaction.ID, err)
	}

	return err
}

func (t *transactionStorage) FindByID(id string) (*domain.Transaction, error) {
	transaction := &domain.Transaction{}

	err := t.DB.QueryRow(
		`SELECT id, agency, account_id, release_date, accounting_date, title, description, income, expense, daily_balance 
		 FROM tb_transactions WHERE id = $1`,
		id,
	).Scan(
		&transaction.ID,
		&transaction.AGENCY,
		&transaction.ACCOUNT_ID,
		&transaction.RELEASE_DATE,
		&transaction.ACCOUNTING_DATE,
		&transaction.TITLE,
		&transaction.DESCRIPTION,
		&transaction.INCOME,
		&transaction.EXPENSE,
		&transaction.DAILY_BALANCE,
	)

	if err != nil {
		log.Printf("Error finding transaction with ID %s: %v", id, err)
		return nil, err
	}

	return transaction, nil
}

func (t *transactionStorage) Delete(id string) error {
	_, err := t.DB.Exec(``, id)

	if err != nil {
		log.Printf("Error deleting user with ID: %s: %v", id, err)
	}

	return err
}

func (t *transactionStorage) ImportTransactions(transactions []*domain.Transaction) error {

	tx, err := t.DB.Begin()

	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return err
	}

	stmt, err := tx.Prepare(``)

	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, transaction := range transactions {
		_, err := stmt.Exec(
			transaction.ID,
			transaction.AGENCY,
			transaction.ACCOUNT_ID,
			transaction.RELEASE_DATE,
			transaction.ACCOUNTING_DATE,
			transaction.TITLE,
			transaction.DESCRIPTION,
			transaction.INCOME,
			transaction.EXPENSE,
			transaction.DAILY_BALANCE,
		)

		if err != nil {
			log.Printf("Error inserting transaction with ID %s: %v", transaction.ID, err)
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
