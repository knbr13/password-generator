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
	fmt.Println("password:", password)
}
