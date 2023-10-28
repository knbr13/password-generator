package main

import (
	"fmt"
	"log"
)

func main() {
	pwd, err := GetPasswordArgs()
	if err != nil {
		log.Fatalln(err)
	}

	charset := GetCharSet(pwd)

	password := BuildPassword(charset, pwd.Length)
	fmt.Printf("%-10s %s\n", "password:", password)
	fmt.Printf("%-10s %s\n", "strength:", PasswordStrength(password))
}
