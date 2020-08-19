package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

type User struct {
	ID          int `json:"id"`
	Email 		string `json:"email"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Password 	string `json:"password"`
}
//HashPassword hashes password from user input
func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword checks password hash and password from user input if they match
func CheckPasswordHash(password, hash string) error {
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
func (u *User) SaveUser(db *sql.DB) (*User, error) {
	var err error

	err = db.QueryRow("insert into users(first_name, last_name, email, password) values ($1, $2, $3, $4) returning id", u.FirstName, u.LastName, u.Email, u.Password).
		Scan(&u.ID)
	if err != nil {
		return &User{}, err
	}
	return u, nil
}


//GetUser returns a user based on email
func (u *User) GetUser(db *sql.DB) (*User, error){
	account := &User{}
	var err  error
	defer func() {
		log.Printf("get user: err %v", err)
	}()
	//log.Printf(u.Email)
	err = db.QueryRow("select first_name, last_name, id, password from users where email = $1", u.Email).
		Scan(&u.FirstName, &u.LastName, &u.ID, &u.Password)
	if err != nil {
		return nil, err
	}

	return account, nil
}

//GetAllUsers returns a list of all the users
func GetAllUsers(db *sql.DB , start, count int) ([]User, error) {
	rows, err := db.Query("select first_name, last_name, email, id, password from users limit $1 offset $2", count, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, u.FirstName, u.LastName, u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}








































































































































































































































































































































































































































































































































































































































































































































































