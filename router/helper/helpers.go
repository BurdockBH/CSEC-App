package helper

import (
	"CSEC-App/config"
	"CSEC-App/statusCodes"
	"CSEC-App/viewmodels"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
	"time"
)

// ValidateToken validates the token
func ValidateToken(tokenString string) (jwt.MapClaims, error) {

	jwtSecret := []byte(config.CFG.JWTConfig.JWTSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// GenerateToken generates a JWT token
func GenerateToken(u *viewmodels.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	jwtSecret := []byte(config.CFG.JWTConfig.JWTSecret)

	claims := token.Claims.(jwt.MapClaims)
	claims["role"] = u.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		println("Failed to sign token:", err)
		return "", err
	}

	return tokenString, nil
}

// HashPassword hashes the password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHashedPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func BaseResponse(w http.ResponseWriter, response []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err := w.Write(response)
	if err != nil {
		println("Error writing response:", err)
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func CheckToken(w *http.ResponseWriter, r *http.Request) jwt.MapClaims {
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		log.Println("Token not found")
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenNotFound,
			Message:    statusCodes.StatusCodes[statusCodes.TokenNotFound],
		})
		BaseResponse(*w, response, http.StatusBadRequest)
		return nil
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	claims, err := ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenValidationFailed,
			Message:    statusCodes.StatusCodes[statusCodes.TokenValidationFailed],
		})
		BaseResponse(*w, response, http.StatusBadRequest)
		return nil
	}

	return claims
}
