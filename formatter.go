// Package msisdn_tools provides a set of helper functions to format a msisdn.
package msisdn_tools

import (
	"errors"
	"regexp"
)

const (
	// CleanerMsisdnFormat is a super clean predefined format.
	CleanerMsisdnFormat = "$1$2$3$4"

	// CleanMsisdnFormat is a clean predefined format.
	CleanMsisdnFormat = "7$1$2$3$4"

	// ShortMsisdnFormat is a short predefined format.
	ShortMsisdnFormat = "$1 $2-$3$4"

	// UsualMsisdnFormat is a usual predefined format.
	UsualMsisdnFormat = "($1) $2-$3$4"

	// FullMsisdnFormat is a fully formatted predefined format.
	FullMsisdnFormat = "+7 ($1) $2-$3$4"
)

var (
	validMsisdn      = regexp.MustCompile(`^7?[69]\d{9}$`)
	notDigitsNorPlus = regexp.MustCompile(`[^\d\+]+`)
	cleanMsisdnParts = regexp.MustCompile(`^(?:\+?7|8)?(\d{3})(\d{3})(\d{2})(\d{2})$`)

	// ErrInvalidMsisdn is an error when formatted msisdn is not valid.
	ErrInvalidMsisdn = errors.New("msisdn is invalid")
)

func cleanMsisdn(msisdn string) (string, error) {
	cleaned := notDigitsNorPlus.ReplaceAllString(msisdn, "")
	formatted := cleanMsisdnParts.ReplaceAllString(cleaned, CleanerMsisdnFormat)

	if !validMsisdn.MatchString(formatted) {
		return "", ErrInvalidMsisdn
	}

	return formatted, nil
}

// IsValidMsisdn determines whether given msisdn is valid.
func IsValidMsisdn(msisdn string) bool {
	if _, err := cleanMsisdn(msisdn); err != nil {
		return false
	}

	return true
}

// FormatMsisdn converts any msisdn to a predefined or custom format pattern.
func FormatMsisdn(msisdn string, pattern string) (string, error) {
	cleaned, err := cleanMsisdn(msisdn)

	if err != nil {
		return "", err
	}

	formatted := cleanMsisdnParts.ReplaceAllString(cleaned, pattern)

	return formatted, nil
}
