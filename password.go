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

func WithLowerCaseChars(b bool) Option {
	return func(p *Password) {
		p.WithLowerCaseChars = b
	}
}

func WithUpperCaseChars(b bool) Option {
	return func(p *Password) {
		p.WithUpperCaseChars = b
	}
}

func WithDigitsChars(b bool) Option {
	return func(p *Password) {
		p.WithDigitsChars = b
	}
}

func WithSpecialChars(b bool) Option {
	return func(p *Password) {
		p.WithSpecialChars = b
	}
}

func OfStrength(s Strength) Option {
	return func(p *Password) {
		p.Strength = s
	}
}
