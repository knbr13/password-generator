package main

import (
	"flag"
	"fmt"
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
