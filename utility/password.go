package utility

import "golang.org/x/crypto/bcrypt"

// HashPassword will hash the password and return the hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	Password := string(bytes)
	return Password, nil
}

// CheckPassword function will check the provided password with users password
func CheckPassword(providedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	return err == nil
}
