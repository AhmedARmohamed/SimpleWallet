package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(100);unique_index" json:"email"`
	FirstName string`gorm:"size:100;not null" 			json:"first_name"`
	LastName string `gorm:"size:100;not null" 			json:"last_name"`
	Password string `gorm:"size:100;not null" 			json:"password"`
}

//HashPassword hashes password from user input
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword checks password hash and password from user input if they match
func CheckPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("Password incorrect")
	}
	return nil
}

//BeforeSave hashes user password
func (u *User) BeforeSave() error {
	password := strings.TrimSpace(u.Password)
	hashedPassword, err := HashPassword(password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

//Prepares strips user input of any white spaces
func (u *User) Prepare() {
	u.Email = strings.TrimSpace(u.Email)
	u.FirstName = strings.TrimSpace(u.LastName)
}

//validate user input
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Email == "" {
			return errors.New("Email is required")
		}

		if u.Password == "" {
			return errors.New("Password is required")
		}

		return nil

	default:
		if u.FirstName == "" {
			return errors.New("FirstName is required")
		}
		if u.LastName == "" {
			return errors.New("lastName is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		return nil
	}
}

//SaveUser adds adds a user to the database
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	//Debug a dingle operation, show detailed log of this operation
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

//GetUser returns a user based on email
func (u *User) GetUser(db *gorm.DB) (*User, error) {
	account := &User{}
	if err := db.Debug().Table("users").Where("email = ?", u.Email).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

//GetAllUsers returns a list of all the user
func(u *User) GetAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	if err := db.Debug().Table("users").Find(&users).Error; err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

































































































































































































































































































































































































































































































































































































































































































































































































































