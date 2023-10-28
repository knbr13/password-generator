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
	WithDigitsChars    bool
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

func WithDigitsChars(b bool) option {
	return func(p *Password) {
		p.WithDigitsChars = b
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
	digitschars := flag.Bool("digitschars", false, "password will contain digitschars characters")
	specialchars := flag.Bool("specialchars", false, "password will contain specialchars characters")
	passwordLength := flag.Uint("passwordLength", 8, "password length")

	if !*lowercase && !*uppercase && !*specialchars && !*digitschars {
		return nil, fmt.Errorf("you must at least choose one set of characters")
	}

	return new(
		WithDigitsChars(*digitschars),
		WithLowerCaseChars(*lowercase),
		WithUpperCaseChars(*uppercase),
		WithSpecialChars(*specialchars),
		OfLength(*passwordLength),
	), nil
}

const (
	LOWER_CASE_CHARS = "abcdefghijklmnopqrstuvwxyz"
	UPPER_CASE_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGITS_CHARS     = "0123456789"
	SPECIAL_CHARS    = "!@#$%^&*()_+[]{}|;':,.<>?/~"
)

func GetCharSet(p *Password) (charset []string) {
	if p.WithUpperCaseChars {
		charset = append(charset, strings.Split(UPPER_CASE_CHARS, "")...)
	}
	if p.WithLowerCaseChars {
		charset = append(charset, strings.Split(LOWER_CASE_CHARS, "")...)
	}
	if p.WithDigitsChars {
		charset = append(charset, strings.Split(DIGITS_CHARS, "")...)
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