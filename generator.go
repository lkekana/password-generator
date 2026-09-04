package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func randomInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return int(n.Int64())
}

func generatePassword(length int, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars bool) ([]byte, error) {
	charset := ""
	if includeUppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if includeLowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if includeNumbers {
		charset += "0123456789"
	}
	if includeSpecialChars {
		charset += "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	}

	if len(charset) == 0 {
		return nil, errors.New("no character sets selected")
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[randomInt(len(charset))]
	}

	return password, nil
}
