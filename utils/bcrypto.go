package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	if err != nil {
		return err
	}
	*password = string(bytes)
	return nil
}

func CheckPassword(dbPass string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
