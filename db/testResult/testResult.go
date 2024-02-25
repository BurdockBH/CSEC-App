package testResult

import (
	"CSEC-App/db"
	"CSEC-App/viewmodels"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func CreateTestResult(testResult *viewmodels.TestResult, claim jwt.MapClaims) error {
	query := "CALL CreateTestResult(?, ?, ?, ?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL CreateTestReesult(%v, %v, %v, %v): %v", testResult.FirstName, testResult.LastName, testResult.Date, testResult.TreatedAt, err)
		return err
	}
	defer st.Close()

	role := claim["role"].(string)

	if role != "lab_technician" {
		log.Printf("User with role %v is not authorized to create test results", role)
		return errors.New("user is not authorized to create test results")
	}

	var created int

	err = st.QueryRow(testResult.FirstName, testResult.LastName, testResult.Date, testResult.TreatedAt).Scan(&created)
	if err != nil {
		log.Printf("Error executing query: CALL CreateTestResult(%v, %v, %v, %v): %v", testResult.FirstName, testResult.LastName, testResult.Date, testResult.TreatedAt, err)
		return err
	}

	if created != 1 {
		log.Printf("Could not create test result with first name %v", testResult.FirstName)
		return err
	}

	return nil
}

func DeleteTestResult(result *viewmodels.TestIdRequest, claim jwt.MapClaims) error {
	query := "CALL DeleteTestResult(?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL DeleteTestResult(%v): %v", result.ID, err)
		return err
	}
	defer st.Close()

	role := claim["role"].(string)

	if role != "lab_technician" {
		log.Printf("User with role %v is not authorized to create test results", role)
		return errors.New("user is not authorized to create test results")
	}

	var deleted int
	err = st.QueryRow(result.ID).Scan(&deleted)
	if err != nil {
		log.Printf("Error executing query: CALL DeleteTestResult(%v): %v", result.ID, err)
		return err
	}

	if deleted != 1 {
		log.Printf("Test result with id %v does not exist", result.ID)
		return fmt.Errorf("test result with id %v does not exist", result.ID)
	}

	return nil
}

func EditTestResult(testResult *viewmodels.EditTestResultRequest, claim jwt.MapClaims) error {
	query := "CALL EditTestResult(?, ?, ?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL EditTestResult(%v, %v): %v", testResult.ID, testResult.Notes, err)
		return err
	}
	defer st.Close()

	role := claim["role"].(string)

	if role != "doctor" {
		log.Printf("User with role %v is not authorized to create test results", role)
		return errors.New("user is not authorized to create test results")
	}

	var edited int
	err = st.QueryRow(testResult.ID, testResult.Type, testResult.Notes).Scan(&edited)
	if err != nil {
		log.Printf("Error executing query: CALL EditTestResult(%v, %v): %v", testResult.ID, testResult.Notes, err)
		return err
	}

	if edited != 1 {
		log.Printf("Test result with id %v does not exist", testResult.ID)
		return fmt.Errorf("product with id %v does not exist", testResult.ID)
	}

	return nil
}

func GetTestResults() ([]viewmodels.TestResult, error) {
	query := "CALL GetTestResults()"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL GetTestResults(): %v", err)
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		log.Printf("Error executing query: CALL GetTestResults(): %v", err)
		return nil, err
	}
	defer rows.Close()

	var testResults []viewmodels.TestResult
	for rows.Next() {
		var testResult viewmodels.TestResult
		var typeValue sql.NullString
		var notesValue sql.NullString

		err = rows.Scan(&testResult.ID, &testResult.FirstName, &testResult.LastName, &typeValue, &notesValue, &testResult.Date, &testResult.TreatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		if typeValue.Valid {
			testResult.Type = typeValue.String
		}

		if notesValue.Valid {
			testResult.Notes = notesValue.String
		}

		testResults = append(testResults, testResult)
	}

	if len(testResults) == 0 {
		log.Println("No results found")
		return nil, errors.New("no result found")
	}

	return testResults, nil
}
