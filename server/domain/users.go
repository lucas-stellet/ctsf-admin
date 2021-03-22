package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base     `valid:"required"`
	Name     string `json:"name" valid:"notnull"`
	Email    string `json:"email" valid:"notnull, email"`
	Password string `json:"password" valid:"notnull"`
}

func (u *User) isValid() error {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) Prepare() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.Password = string(password)

	err = user.isValid()

	if err != nil {
		return err
	}

	return nil
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func NewUser(name, email, password string) (*User, error) {
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.isValid()

	if err != nil {
		return nil, err
	}

	return user, nil
}
