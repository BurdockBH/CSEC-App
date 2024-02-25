package viewmodels

// BaseResponse is the response model
type BaseResponse struct {
	StatusCode int64  `json:"status_code"`
	Message    string `json:"message"`
}

// LoginResponse is the response model for login
type LoginResponse struct {
	BaseResponse
	AccessToken string `json:"token"`
}

// UserList is the response model for user
type UserList struct {
	BaseResponse
	Users []User `json:"users"`
}

type TestResultList struct {
	BaseResponse
	TestResults []TestResult `json:"test_results"`
}
