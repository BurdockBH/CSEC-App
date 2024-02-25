package viewmodels

import (
	"errors"
	"regexp"
)

type TestResult struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"Type"`
	Notes     string `json:"notes"`
	Date      int64  `json:"date"`
	TreatedAt string `json:"treated_at"`
}

type TestIdRequest struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type EditTestResultRequest struct {
	ID    string `json:"id"`
	Type  string `json:"Type"`
	Notes string `json:"notes"`
}

func (u *TestResult) Validate() error {
	if u.FirstName == "" || len(u.FirstName) > 50 || len(u.FirstName) < 2 || regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(u.FirstName) == false {
		return errors.New("invalid first name")
	} else if u.LastName == "" || len(u.LastName) > 50 || len(u.LastName) < 2 || regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(u.LastName) == false {
		return errors.New("invalid last name")
	} else if len(u.Type) > 50 {
		return errors.New("invalid Type")
	} else if len(u.Notes) > 2000 {
		return errors.New("invalid notes")
	} else if u.Date == 0 {
		return errors.New("invalid date")
	} else if u.TreatedAt == "" || len(u.TreatedAt) > 50 || len(u.TreatedAt) < 2 {
		return errors.New("invalid treatedAt")
	}
	return nil
}

func (i *TestIdRequest) ValidateItemIdRequest() error {
	if len(i.ID) == 0 {
		return errors.New("id must not be null")
	}
	return nil
}

func (e *EditTestResultRequest) ValidateEditTestResultRequest() error {
	if e.ID == "" {
		return errors.New("id must not be null")
	}
	if len(e.Type) > 50 {
		return errors.New("invalid Type")
	}
	if len(e.Notes) > 2000 {
		return errors.New("invalid notes")
	}
	return nil
}
