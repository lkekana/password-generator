package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"runtime"

	"github.com/fatih/color"
)

func randomInt(max *big.Int) int {
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return int(n.Int64())
}

// ensure password has all the required character types based on the user's selection
func meetsRequirements(pwd []byte, upper, lower, num, special bool) bool {
	var hasUpper, hasLower, hasNum, hasSpecial bool
	for _, c := range pwd {
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= '0' && c <= '9':
			hasNum = true
		default:
			hasSpecial = true
		}
	}
	// only returns true if the password meets all the requirements specified by the user
	return (!upper || hasUpper) && (!lower || hasLower) && (!num || hasNum) && (!special || hasSpecial)
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

	// ensure that the password length is long enough for the required character types
	// (prevents infinite loops in rejection sampling below)
	requiredCount := 0
	if includeUppercase {
		requiredCount++
	}
	if includeLowercase {
		requiredCount++
	}
	if includeNumbers {
		requiredCount++
	}
	if includeSpecialChars {
		requiredCount++
	}

	if length < requiredCount {
		return nil, fmt.Errorf("password length is too short for the selected character requirements. Minimum length required: %d", requiredCount)
	}

	password := make([]byte, length)
	maxBig := big.NewInt(int64(len(charset)))

	for {
		for i := range password {
			password[i] = charset[randomInt(maxBig)]
		}

		if meetsRequirements(password, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars) {
			return password, nil
		} else {
			if debug {
				color.Yellow("Generated password did not meet requirements, regenerating...")
			}
		}
	}
}

func zeroOutPassword(password []byte) {
	for i := range password {
		password[i] = 0
	}
	// ensures the password will only be dealt with by the garbage collector after this function has completed
	// will ensure the function is not optimized away by the compiler and is actually zeroed out in memory
	runtime.KeepAlive(password)
}