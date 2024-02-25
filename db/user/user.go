package user

import (
	"CSEC-App/db"
	"CSEC-App/router/helper"
	"CSEC-App/viewmodels"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// Database queries and logic for user

// RegisterUser registers a new user
func RegisterUser(u *viewmodels.User) error {

	hashedPassword, err := helper.HashPassword(u.Password)
	if err != nil {
		println("Error hashing password:", err)
		return err
	}

	query := "CALL RegisterUser(?, ?, ?, ?, ?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf(`Error preparing query "CALL RegisterUser(%v, %v, %v, %v, %v): %v"`, u.FirstName, u.LastName, u.Email, hashedPassword, u.Role, err)
		return err
	}
	defer st.Close()

	var created int
	err = st.QueryRow(u.FirstName, u.LastName, u.Email, hashedPassword, u.Role).Scan(&created)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	if created != 1 {
		log.Printf("User with email: %v already exists", u.Email)
		return fmt.Errorf("user with email %v already exists", u.Email)
	}

	return nil
}

// LoginUser logs in a user, it checks if the user exists and if the password matches
func LoginUser(u *viewmodels.UserLoginRequest) (error, *viewmodels.User) {
	var userInfo = viewmodels.User{}

	query := "CALL LoginUser(?)"

	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf(`Error preparing query "CALL LoginUser(%v)": %v`, u.Email, err)
		return err, nil
	}

	rows, err := st.Query(u.Email)
	if err != nil {
		log.Println("Error executing query:", err)
		return err, nil
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userInfo.ID, &userInfo.FirstName, &userInfo.LastName, &userInfo.Email, &userInfo.Password, &userInfo.Role)
		if err != nil {
			log.Println("Error scanning row:", err)
			return err, nil
		}
	}

	if userInfo.ID == "" {
		log.Printf("User with email %v does not exist", u.Email)
		return fmt.Errorf("user with email %v does not exist", u.Email), nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(u.Password))
	if err != nil {
		log.Println("Error comparing password:", err)
		return errors.New("error comparing password"), nil
	}
	return nil, &userInfo
}

// DeleteUser deletes a user from the database
func DeleteUser(u *viewmodels.UserLoginRequest) error {
	passwordQuery := "CALL LoginUser(?)"
	var password string
	err := db.DB.QueryRow(passwordQuery, u.Email).Scan(&password)
	if err != nil {
		log.Println("User does not exist:", err)
		return fmt.Errorf("user %v does not exist", u.Email)
	}

	err = helper.CompareHashedPassword(password, u.Password)
	if err != nil {
		log.Println("Error comparing password:", err)
		return errors.New("error comparing password")
	}

	query := "CALL DeleteUser(?)"

	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf(`Error preparing query "CALL DeleteUser(%v)": %v`, u.Email, err)
		return err
	}

	var deleted int
	err = st.QueryRow(u.Email).Scan(&deleted)
	if err != nil {
		log.Printf("Failed to delete user with email %v. error is %v \n ", u.Email, err)
		return err
	}

	if deleted != 1 {
		log.Printf("Couldn't delete %v. No rows affected\n", u.Email)
		return fmt.Errorf("couldn't delete user %v. No rows affected", u.Email)
	}

	return nil
}

// EditUser edits a user's information
func EditUser(u *viewmodels.User) error {
	query := "CALL EditUser(?, ?, ?, ?, ?)"

	st, err := db.DB.Prepare(query)
	if err != nil {
		//		log.Printf(`Error preparing query "CALL EditUser(%v, %v, %v, %v, %v": %v`, u.Name, u.Email, u.Password, u.Phone, time.Now().Unix(), err)
		return err
	}

	hashedPassword, err := helper.HashPassword(u.Password)
	if err != nil {
		log.Println("Failed to hash password", err)
		return err
	}

	var updated int
	err = st.QueryRow(u.FirstName, u.Email, hashedPassword, u.LastName, time.Now().Unix()).Scan(&updated)
	if err != nil {
		log.Printf("Failed to update user with email %v. Error: %v\n\n\n", u.Email, err)
		return err
	}

	if updated == -1 {
		log.Printf("User with email %v does not exist", u.Email)
		return fmt.Errorf("user with email %v does not exist", u.Email)
	} else if updated == -2 {
		//		log.Printf("User with phone number: %v already exists", u.Phone)
		//		return fmt.Errorf("user with phone number: %v already exists", u.Phone)
	}

	return nil
}

func GetUsers(u *viewmodels.User) ([]viewmodels.User, error) {
	query := "CALL GetUsers(?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf(`Error preparing query "CALL GetUsersByDetails(%v, %v, %v)": %v`, u.FirstName, u.Email, u.LastName, err)
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(u.ID)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var users []viewmodels.User
	for rows.Next() {
		var user viewmodels.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role)
		if err != nil {
			log.Println("Error scanning row:", err)
			users = append(users, viewmodels.User{})
		}
		users = append(users, user)
	}

	if users == nil {
		log.Println("No users found")
		return nil, errors.New("no users found")
	}

	return users, nil
}
