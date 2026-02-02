package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
}

func ComparePassword(password, hashed string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashed), []byte(password),
	)
}
