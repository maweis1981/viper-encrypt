package main

import "fmt"

func encryptString(plaintext string, key string) (string, error) {
	password, err := encrypt(plaintext, key)
	if err != nil {
		fmt.Println("Error encrypting password: ", err)
	}
	fmt.Println(decrypt(password, key))
	fmt.Println("------------------")
	return password, err
}
