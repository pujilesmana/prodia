package utils

import "golang.org/x/crypto/bcrypt"

func Hashing(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CompareHash(input, target string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(input), []byte(target)); err != nil {
		return err
	}

	return nil
}
