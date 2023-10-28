package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

type Strength uint8

const (
	WEAK Strength = iota
	MEDIUM
	STRONG
)

type Password struct {
	WithLowerCaseChars bool
	WithUpperCaseChars bool
	WithSpecialChars   bool
	WithDigitsChars    bool
	Strength           Strength
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

func OfStrength(s Strength) option {
	return func(p *Password) {
		p.Strength = s
	}
}

func GetPasswordArgs() (*Password, error) {
	lowercase := flag.Bool("lowercase", true, "password will contain lowercase characters")
	uppercase := flag.Bool("uppercase", true, "password will contain uppercase characters")
	digitschars := flag.Bool("digitschars", false, "password will contain digitschars characters")
	specialchars := flag.Bool("specialchars", false, "password will contain specialchars characters")

	if !*lowercase && !*uppercase && !*specialchars && !*digitschars {
		return nil, fmt.Errorf("password: you must at least choose one set of characters")
	}

	return new(
		WithDigitsChars(*digitschars),
		WithLowerCaseChars(*lowercase),
		WithUpperCaseChars(*uppercase),
		WithSpecialChars(*specialchars),
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

func BuildPassword(charset []string, length int) string {
	var s string
	for i := 0; i < length; i++ {
		s += charset[rand.Intn(len(charset))]
	}
	return s
}
