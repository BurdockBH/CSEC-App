package testResult

import (
	"CSEC-App/db/testResult"
	"CSEC-App/router/helper"
	"CSEC-App/statusCodes"
	"CSEC-App/viewmodels"
	"encoding/json"
	"log"
	"net/http"
)

func CreateTestResult(w http.ResponseWriter, r *http.Request) {
	var t viewmodels.TestResult

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	claims := helper.CheckToken(&w, r)
	if claims == nil {
		return
	}

	err = t.Validate()
	if err != nil {
		log.Println("Failed to validate test result:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateTestResult,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateTestResult],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = testResult.CreateTestResult(&t, claims)
	if err != nil {
		log.Println("Failed to create result:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToCreateTestResult,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToCreateTestResult] + ": " + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusUnauthorized)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyCreatedTestResult,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyCreatedTestResult] + ": " + t.FirstName,
	})
	log.Println("Successfully created test result:", t.FirstName)
	helper.BaseResponse(w, response, http.StatusOK)
}

func DeleteTestResult(w http.ResponseWriter, r *http.Request) {
	var idRequest viewmodels.TestIdRequest

	err := json.NewDecoder(r.Body).Decode(&idRequest)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	claims := helper.CheckToken(&w, r)
	if claims == nil {
		return
	}

	err = idRequest.ValidateItemIdRequest()
	if err != nil {
		log.Println("Failed to validate test result:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateTestResult,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateTestResult],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = testResult.DeleteTestResult(&idRequest, claims)
	if err != nil {
		log.Println("Failed to delete test result:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDeleteTestResult,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDeleteTestResult] + ": " + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyDeletedTestResult,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyDeletedTestResult] + ": " + idRequest.ID,
	})
	log.Println("Successfully Deleted test result:", idRequest.ID)
	helper.BaseResponse(w, response, http.StatusOK)
}

func EditTestResult(w http.ResponseWriter, r *http.Request) {
	var result viewmodels.EditTestResultRequest

	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	claims := helper.CheckToken(&w, r)
	if claims == nil {
		return
	}

	err = result.ValidateEditTestResultRequest()
	if err != nil {
		log.Println("Failed to validate test result:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateTestResult,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateTestResult],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = testResult.EditTestResult(&result, claims)
	if err != nil {
		log.Println("Failed to edit test result:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToEditTestResult,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToEditTestResult] + ": " + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusUnauthorized)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyEditedTestResult,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyEditedTestResult] + ": " + result.ID,
	})
	log.Println("Successfully edited test result:", result.ID)
	helper.BaseResponse(w, response, http.StatusOK)
}

func GetTestResults(w http.ResponseWriter, r *http.Request) {
	claims := helper.CheckToken(&w, r)
	if claims == nil {
		return
	}

	testResults, err := testResult.GetTestResults()
	if err != nil {
		log.Println("Failed to register user:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToGetTestResults,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToGetTestResults] + ": " + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.TestResultList{
		BaseResponse: viewmodels.BaseResponse{
			StatusCode: statusCodes.SuccesfullyGetTestResults,
			Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyGetTestResults],
		},
		TestResults: testResults,
	})
	log.Println("Successfully retrieved test results")
	helper.BaseResponse(w, response, http.StatusOK)
}
