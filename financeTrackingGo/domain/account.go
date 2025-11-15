package domain

type Account struct {
	ID           string  `json:"id"`
	ACCOUNT_TYPE string  `json:"account_type"`
	ACTIVE       bool    `json:"active"`
	BALANCE      float64 `json:"balance"`
	BANK_NAME    string  `json:"bank_name"`
	CREATED_AT   string  `json:"created_at"`
	UPDATED_AT   string  `json:"updated_at"`
	USER_ID      string  `json:"user_id"`
}

type AccountService interface {
	Post(account *Account) (*Account, error)
	Put(Account *Account) error
	Delete(id string) error
	GetByID(id string) (*Account, error)
}

type AccountStorage interface {
	Insert(account *Account) (*Account, error)
	Update(Account *Account) error
	Delete(id string) error
	FindByID(id string) (*Account, error)
}

func NewAccount(
	id string,
	accountType string,
	active bool,
	balance float64,
	bankName string,
	createdAt string,
	updatedAt string,
	userID string,
) *Account {
	return &Account{
		ID:           id,
		ACCOUNT_TYPE: accountType,
		ACTIVE:       active,
		BALANCE:      balance,
		BANK_NAME:    bankName,
		CREATED_AT:   createdAt,
		UPDATED_AT:   updatedAt,
		USER_ID:      userID,
	}
}
