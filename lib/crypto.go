package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

//Hash implements root.Hash
type Hash struct{}

//Generate a salted hash for the input string
func Generate(s string) string {
	if s == "" {
		return ""
	}
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	hash := string(hashedBytes[:])
	return hash
}

//Compare string to generated hash
func Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
