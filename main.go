package main

import "log"

func main() {
	pwd, err := GetPasswordArgs()
	if err != nil {
		log.Fatalln(err)
	}

	charset := GetCharSet(pwd)

	BuildPassword(charset, pwd.Length)
}
