package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword := string(p)
	return hashedPassword, err
}

func CompareHashedPassword(dbPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
}
