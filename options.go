package main

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
