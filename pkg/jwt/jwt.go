package jwt

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = "yourSigningKey"

func NewToken(user_id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func RetrieveUserIDFromRequest(r *http.Request) (int, error) {
	return retrieveClaimFromRequest(r, "uid")
}

func retrieveClaimFromRequest(r *http.Request, claimKey string) (int, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("authorization header is missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return 0, errors.New("invalid authorization header format")
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token or claims")
	}

	claimValue, ok := claims[claimKey].(float64)
	if !ok {
		return 0, errors.New(claimKey + " not found or invalid type")
	}

	return int(claimValue), nil
}
