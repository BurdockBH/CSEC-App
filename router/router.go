package router

import (
	middlewares "CSEC-App/router/middleware"
	"CSEC-App/service/testResult"
	"CSEC-App/service/user"
	"net/http"
)

// InitializeRouter initializes the router
func InitializeRouter() *http.ServeMux {
	router := http.NewServeMux()

	// User routes
	router.HandleFunc("/user/register", middlewares.Chain(middlewares.Post)(user.RegisterUser))
	router.HandleFunc("/user/login", middlewares.Chain(middlewares.Post)(user.LoginUser))
	router.HandleFunc("/user/delete", middlewares.Chain(middlewares.Delete)(user.DeleteUser))
	router.HandleFunc("/user/edit", middlewares.Chain(middlewares.Put)(user.EditUser))
	router.HandleFunc("/users/get", middlewares.Chain(middlewares.Get)(user.GetUsers))

	// Test Results routes
	router.HandleFunc("/test-result/create", middlewares.Chain(middlewares.Post)(testResult.CreateTestResult))
	router.HandleFunc("/test-result/delete", middlewares.Chain(middlewares.Delete)(testResult.DeleteTestResult))
	router.HandleFunc("/test-result/get", middlewares.Chain(middlewares.Get)(testResult.GetTestResults))
	router.HandleFunc("/test-result/edit", middlewares.Chain(middlewares.Put)(testResult.EditTestResult))

	return router
}
