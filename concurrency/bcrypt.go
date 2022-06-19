package main

import (
	"fmt"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// in this APPLICATION, I will use 3 CORES (CPUs) to hash 3 passwords simultaneously

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go pwd1()
	go pwd2()
	pwd3()
	wg.Wait()
}

func pwd1()  {
	// collects "new password"
	password := "1234"
	byte_password := []byte(password)

	// "hash" the registered password
	hash_password, err := bcrypt.GenerateFromPassword(byte_password, 16)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hash_password)
	wg.Done()
}

func pwd2()  {
	// collects "new password"
	password2 := "1234"
	byte_password2 := []byte(password2)

	// "hash" the registered password
	hash_password2, err2 := bcrypt.GenerateFromPassword(byte_password2, 15)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(hash_password2)
	wg.Done()
}

func pwd3()  {
	// collects "new password"
	password3 := "1234"
	byte_password3 := []byte(password3)

	// "hash" the registered password
	hash_password3, err3 := bcrypt.GenerateFromPassword(byte_password3, 14)
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(hash_password3)
}