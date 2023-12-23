package customer_management

import (
	"golang.org/x/crypto/bcrypt"
)

func createCustomer(Username string, Password string) error {
	// Creates new customer account.
}

func deleteCustomer(Username string, Password string) error {
	// Deletes user account.
}

func loadCustomer(Username string, Password string) (Customer, error) {
	// Loads customer from database.
}

func saveCustomer(c Customer) {
	// Writes Cx to database.
}

func verifyEmail(Email string) {
	// Create string, send email to email with known-good domain suffix,
	// confirm the checksum matches the other sum.
}

func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
