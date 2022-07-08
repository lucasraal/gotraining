package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	// collects "new password"
	password := "1234"
	byte_password := []byte(password)

	// "hash" the registered password
	hash_password, err := bcrypt.GenerateFromPassword(byte_password, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	// asks for "password" when logging in to the server
	login_password := "12344"
	byte_login_password := []byte(login_password)

	// checks if the "entered password" (after hashing it) is equal to the registered "hashed" password
	err = bcrypt.CompareHashAndPassword(hash_password, byte_login_password)
	if err != nil {
		fmt.Println("YOU CANNOT LOG IN\nerror:",err)
		// function "main" terminates after executing the return statement
		return
	}
	
	fmt.Println("YOU HAVE SUCCESSFULLY LOGGED IN")
}