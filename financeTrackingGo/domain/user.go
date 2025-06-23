package domain

type User struct {
	ID         string `json:"id"`
	NAME       string `json:"name"`
	EMAIL      string `json:"email"`
	PASSWORD   string `json:"password"`
	CPF        string `json:"cpf"`
	CREATED_AT string `json:"created_at"`
	UPDATED_AT string `json:"updated_at"`
	ACITVE     bool   `json:"active"`
}

type UserService interface {
	GetByCpf(cpf string) (*User, error)
	GetByID(id string) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) error
	Delete(id string) error
}

type UserStorage interface {
	Insert(user *User) (*User, error)
	Update(user *User) error
	Delete(id string) error
	FindByID(id string) (*User, error)
	FindByCpf(cpf string) (*User, error)
}

func NewUser(
	id string,
	cpf string,
	name string,
	email string,
	password string,
	active bool,
	createdAt string,
	updatedAt string,
) *User {
	return &User{
		CPF:        cpf,
		NAME:       name,
		EMAIL:      email,
		PASSWORD:   password,
		ACITVE:     active,
		CREATED_AT: createdAt,
		UPDATED_AT: updatedAt,
		ID:         id,
	}
}
