package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Model  interface{}
	Role   string
	UserId int
	jwt.RegisteredClaims
}

func CreateToken(secretKey, role string, userId int, expireAfter time.Duration, claim interface{}) (tokenString string, customClaims interface{}, err error) {

	// Create the Custom Claims
	claims := &CustomClaims{
		claim,
		role,
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireAfter)), // Token expires in 24 hours
		},
	}

	// Generate token based on claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Retrieving token string
	tokenString, err = token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error occured while creating token:", err)
		return "", nil, err
	}

	return tokenString, claims.Model, nil
}

// Validate Token
func IsValidToken(secretKey, tokenString string) (bool, interface{}) {

	// Parse jwt token with custom claims
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	// Check if token is valid
	if err != nil || !token.Valid {
		fmt.Println("Error occured while parsing token:", err)
		return false, nil
	}

	// Assign parsed data from token to calims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

		// Check if token is expired
		if claims.ExpiresAt.Before(time.Now()) {
			fmt.Println("token expired")
			return false, nil
		}

		return true, claims

	} else {
		fmt.Println("Error occured while decoding token:", err)
		return false, nil
	}
}
