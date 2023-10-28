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

type Option func(*Password)

func New(options ...Option) *Password {
	p := &Password{
		WithLowerCaseChars: true,
		WithUpperCaseChars: true,
	}
	for _, option := range options {
		option(p)
	}
	return p
}
