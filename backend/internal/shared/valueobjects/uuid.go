package valueobjects

import (
	"errors"
	"regexp"
	"strings"
)

// UUID represents a UUID value object
type UUID struct {
	value string
}

// NewUUID creates a new UUID value object
func NewUUID() UUID {
	// TODO: Generate actual UUID using uuid package
	return UUID{value: "generated-uuid"}
}

// String returns the string representation of UUID
func (u UUID) String() string {
	return u.value
}

// Email represents an email value object
type Email struct {
	value string
}

// NewEmail creates a new Email value object
func NewEmail(email string) (Email, error) {
	if !isValidEmail(email) {
		return Email{}, errors.New("invalid email format")
	}
	return Email{value: strings.ToLower(email)}, nil
}

// String returns the string representation of Email
func (e Email) String() string {
	return e.value
}

// isValidEmail validates email format
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// CountryCode represents a country code value object (ISO 3166-1 alpha-3)
type CountryCode struct {
	value string
}

// NewCountryCode creates a new CountryCode value object
func NewCountryCode(code string) (CountryCode, error) {
	if !isValidCountryCode(code) {
		return CountryCode{}, errors.New("invalid country code format")
	}
	return CountryCode{value: strings.ToUpper(code)}, nil
}

// String returns the string representation of CountryCode
func (c CountryCode) String() string {
	return c.value
}

// isValidCountryCode validates country code format (3 uppercase letters)
func isValidCountryCode(code string) bool {
	countryCodeRegex := regexp.MustCompile(`^[A-Z]{3}$`)
	return countryCodeRegex.MatchString(code)
}
