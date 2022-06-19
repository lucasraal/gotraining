package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "1234"
	byte_password := []byte(password)
	fmt.Println(password, byte_password)

	// GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hash_password, err := bcrypt.GenerateFromPassword(byte_password, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hash_password)
	fmt.Println(string(hash_password))
}