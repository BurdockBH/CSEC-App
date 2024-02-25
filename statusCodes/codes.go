package statusCodes

const (
	UserDoesNotExist             = 1
	FailedToDecodeRequestBody    = 2
	FailedToValidateUser         = 3
	FailedToCreateUser           = 4
	FailedToUpdateUser           = 5
	FailedToValidateLogin        = 6
	FailedToLoginUser            = 7
	FailedToGenerateToken        = 8
	FailedToMarshalJSON          = 9
	FailedToWriteResponse        = 10
	TokenNotFound                = 11
	TokenValidationFailed        = 12
	InvalidClaims                = 13
	FailedToDeleteUser           = 14
	FailedToFetchUsers           = 15
	SuccesfullyCreatedUser       = 16
	SuccesfullyDeletedUser       = 17
	SuccesfullyFetchedUsers      = 18
	SuccesfullyUpdatedUser       = 19
	SuccesfullyLoggedInUser      = 20
	FailedToHashID               = 21
	FailedToValidateTestResult   = 22
	FailedToCreateTestResult     = 23
	SuccesfullyCreatedTestResult = 24
	FailedToEditTestResult       = 25
	SuccesfullyEditedTestResult  = 26
	FailedToDeleteTestResult     = 27
	SuccesfullyDeletedTestResult = 28
	SuccesfullyGetTestResults    = 29
	FailedToGetTestResults       = 30
)

var StatusCodes = map[int64]string{
	UserDoesNotExist:             "User does not exist",
	FailedToDecodeRequestBody:    "Failed to decode request body",
	FailedToValidateUser:         "Failed to validate user",
	FailedToCreateUser:           "Failed to create user",
	FailedToUpdateUser:           "Failed to update user",
	FailedToValidateLogin:        "Failed to validate login",
	FailedToLoginUser:            "Failed to login user",
	FailedToGenerateToken:        "Failed to generate token",
	FailedToMarshalJSON:          "Failed to marshal json",
	FailedToWriteResponse:        "Failed to write response",
	TokenNotFound:                "Token not found",
	TokenValidationFailed:        "Token validation failed",
	InvalidClaims:                "Invalid claims",
	FailedToDeleteUser:           "Failed to delete user",
	FailedToFetchUsers:           "Failed to fetch users",
	SuccesfullyCreatedUser:       "User created successfully!",
	SuccesfullyDeletedUser:       "User deleted successfully!",
	SuccesfullyFetchedUsers:      "Users fetched successfully!",
	SuccesfullyUpdatedUser:       "User updated successfully!",
	SuccesfullyLoggedInUser:      "User logged in successfully!",
	FailedToHashID:               "Failed to hash ID",
	FailedToValidateTestResult:   "Failed to validate test result",
	FailedToCreateTestResult:     "Failed to create test result",
	SuccesfullyCreatedTestResult: "Test result created successfully!",
	FailedToEditTestResult:       "Failed to edit test result",
	SuccesfullyEditedTestResult:  "Test result edited successfully!",
	FailedToDeleteTestResult:     "Failed to delete test result",
	SuccesfullyDeletedTestResult: "Test result deleted successfully!",
	SuccesfullyGetTestResults:    "Test results fetched successfully!",
	FailedToGetTestResults:       "Failed to fetch test results",
}
