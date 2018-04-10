package msisdn_tools

import (
	"testing"
)

func TestFormatMsisdn(t *testing.T) {
	formatTests := []struct {
		Msisdn   string
		Format   string
		Expected string
	}{
		{"9955053777", CleanerMsisdnFormat, "9955053777"},
		{"9955053777", CleanMsisdnFormat, "79955053777"},
		{"9955053777", ShortMsisdnFormat, "995 505-3777"},
		{"9955053777", UsualMsisdnFormat, "(995) 505-3777"},
		{"9955053777", FullMsisdnFormat, "+7 (995) 505-3777"},
	}

	for _, formatTest := range formatTests {
		formatTest := formatTest
		t.Run(formatTest.Format, func(t *testing.T) {
			t.Parallel()
			formatted, err := FormatMsisdn(formatTest.Msisdn, formatTest.Format)

			if err != nil {
				t.Errorf("error while formatting %s to %s: %s",
					formatTest.Msisdn, formatTest.Expected, err.Error())
			}

			if formatTest.Expected != formatted {
				t.Errorf("invalid format result. expected: %s, got: %s", formatTest.Expected, formatted)
			}
		})
	}
}

func TestFormatMsisdnInvalid(t *testing.T) {
	_, err := FormatMsisdn("9912323", CleanerMsisdnFormat)

	if err != ErrInvalidMsisdn {
		t.Errorf("invalid error type. expected: %s, got: %s", ErrInvalidMsisdn.Error(), err.Error())
	}
}
