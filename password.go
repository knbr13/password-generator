package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

type Password struct {
	WithLowerCaseChars bool
	WithUpperCaseChars bool
	WithSpecialChars   bool
	WithDigits         bool
	Length             uint
}

type option func(*Password)

func new(options ...option) *Password {
	p := &Password{
		WithLowerCaseChars: true,
		WithUpperCaseChars: true,
	}
	for _, option := range options {
		option(p)
	}
	return p
}

func WithLowerCaseChars(b bool) option {
	return func(p *Password) {
		p.WithLowerCaseChars = b
	}
}

func WithUpperCaseChars(b bool) option {
	return func(p *Password) {
		p.WithUpperCaseChars = b
	}
}

func WithDigits(b bool) option {
	return func(p *Password) {
		p.WithDigits = b
	}
}

func WithSpecialChars(b bool) option {
	return func(p *Password) {
		p.WithSpecialChars = b
	}
}

func OfLength(s uint) option {
	return func(p *Password) {
		p.Length = s
	}
}

func GetPasswordArgs() (*Password, error) {
	lowercase := flag.Bool("lowercase", true, "password will contain lowercase characters")
	uppercase := flag.Bool("uppercase", true, "password will contain uppercase characters")
	digitschars := flag.Bool("digits", false, "password will contain digitschars characters")
	specialchars := flag.Bool("specialchars", false, "password will contain specialchars characters")
	passwordLength := flag.Uint("length", 8, "password length")
	flag.Parse()

	if !*lowercase && !*uppercase && !*specialchars && !*digitschars {
		return nil, fmt.Errorf("you must at least choose one set of characters")
	}

	return new(
		WithDigits(*digitschars),
		WithLowerCaseChars(*lowercase),
		WithUpperCaseChars(*uppercase),
		WithSpecialChars(*specialchars),
		OfLength(*passwordLength),
	), nil
}

const (
	LOWER_CASE_CHARS = "abcdefghijklmnopqrstuvwxyz"
	UPPER_CASE_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SPECIAL_CHARS    = "!@#$%^&*()_+[]{}|;':,.<>?/~"
	DIGITS           = "0123456789"
)

func GetCharSet(p *Password) (charset []string) {
	if p.WithUpperCaseChars {
		charset = append(charset, strings.Split(UPPER_CASE_CHARS, "")...)
	}
	if p.WithLowerCaseChars {
		charset = append(charset, strings.Split(LOWER_CASE_CHARS, "")...)
	}
	if p.WithDigits {
		charset = append(charset, strings.Split(DIGITS, "")...)
	}
	if p.WithSpecialChars {
		charset = append(charset, strings.Split(SPECIAL_CHARS, "")...)
	}
	return
}

func BuildPassword(charset []string, length uint) (s string) {
	var i uint
	for i = 0; i < length; i++ {
		s += charset[rand.Intn(len(charset))]
	}
	return s
}

func PasswordStrength(password string) string {
	if len(password) < 8 {
		return "Weak"
	}

	var hasLowerCase, hasUpperCase, hasDigit, hasSpecial bool

	for _, char := range password {
		if 'a' <= char && char <= 'z' {
			hasLowerCase = true
		} else if 'A' <= char && char <= 'Z' {
			hasUpperCase = true
		} else if '0' <= char && char <= '9' {
			hasDigit = true
		} else {
			hasSpecial = true
		}
	}

	typesCount := 0
	if hasLowerCase {
		typesCount++
	}
	if hasUpperCase {
		typesCount++
	}
	if hasDigit {
		typesCount++
	}
	if hasSpecial {
		typesCount++
	}

	if typesCount >= 3 {
		return "Strong"
	} else if typesCount == 2 {
		return "Moderate"
	}

	return "Weak"
}
