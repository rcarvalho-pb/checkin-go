package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rcarvalho-pb/checkin-go/internal/config"
	"github.com/rcarvalho-pb/checkin-go/internal/participant"
	participant_role "github.com/rcarvalho-pb/checkin-go/internal/participant/roles"
)

type Claims struct {
	ID    int
	Name  string
	Email string
	Role  participant_role.Role
	jwt.RegisteredClaims
}

func NewJwtToken(p *participant.Participant) (string, error) {
	claims := Claims{
		ID:    p.ID,
		Name:  p.Name,
		Email: p.Email,
		Role:  p.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "participant-login",
			Subject:   "login jwt token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getValidationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(config.Secret), nil
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	if tokenString == "" {
		return fmt.Errorf("invalid token")
	}
	token, err := jwt.Parse(tokenString, getValidationKey)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	if _, ok := token.Claims.(*Claims); !ok || !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func ParseJwtTokenWithClaims(tokenString string) (*Claims, error) {
	var claims *Claims
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, getValidationKey)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
