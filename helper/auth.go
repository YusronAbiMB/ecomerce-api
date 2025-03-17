package helper

import (
	"fmt"
	"net"
	"net/mail"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetUserIDFromCookie(ctx *gin.Context) (uint, error) {

	cookie, err := ctx.Cookie("token")
	if err != nil {
		return 0, fmt.Errorf("Authorization token is missing")
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, fmt.Errorf("Invalid token")
	}
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("Failed to parse token claims")
	}
	userIDFloat, ok := (*claims)["id"].(float64)
	if !ok || userIDFloat == 0 {
		return 0, fmt.Errorf("Missing or invalid user ID in token")
	}
	userID := uint(userIDFloat)
	return userID, nil
}

func GetRoleFromToken(ctx *gin.Context) (string, error) {

	cookie, err := ctx.Cookie("token")
	if err != nil {
		return "", fmt.Errorf("Authorization token is missing")
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", fmt.Errorf("Invalid token")
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("Failed to parse token claims")
	}

	role, ok := (*claims)["role"].(string)
	if !ok || role == "" {
		return "", fmt.Errorf("Missing or invalid role in token")
	}
	return role, nil
}

func MatchesHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	_, err = net.LookupMX(domain)
	if err != nil {
		return false
	}
	return true
}

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	reUppercase := regexp.MustCompile(`[A-Z]`)
	reDigit := regexp.MustCompile(`\d`)
	reSymbol := regexp.MustCompile(`[\W_]`)
	return reUppercase.MatchString(password) && reDigit.MatchString(password) && reSymbol.MatchString(password)
}
