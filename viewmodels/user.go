package viewmodels

import (
	"errors"
	"regexp"
)

// User is the user model
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type UserLoginRequest struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate validates the user
func (u *User) Validate() error {
	if regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email) == false {
		return errors.New("invalid email")
	} else if regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(u.FirstName) == false {
		return errors.New("invalid name")
	} else if regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(u.LastName) == false {
		return errors.New("invalid phone")
	} else if len(u.Password) > 20 || len(u.Password) < 8 {
		return errors.New("invalid password")
	}
	return nil
}

// ValidateLogin validates the user login credentials
func (u *UserLoginRequest) ValidateLogin() error {
	if regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email) == false {
		return errors.New("invalid email")
	} else if len(u.Password) > 20 || len(u.Password) < 8 {
		return errors.New("invalid password")
	}
	return nil
}
