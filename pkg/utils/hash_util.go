package utils

import "golang.org/x/crypto/bcrypt"

var defaultCost = 14

// HashPassword is a function for hashing password using bcrypt
func HashPassword(password string, cost int) (string, error) {
	if cost == 0 {
		cost = defaultCost
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// CheckPasswordHash is a function for comparing password with hash
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
