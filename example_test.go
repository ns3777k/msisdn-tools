package msisdn_tools

import (
	"fmt"
)

func ExampleFormatMsisdn() {
	f, _ := FormatMsisdn("9955053777", FullMsisdnFormat)
	fmt.Println(f)
	// Output: +7 (995) 505-3777
}

func ExampleFormatMsisdn_fail() {
	_, err := FormatMsisdn("8394462", FullMsisdnFormat)
	fmt.Println(err)
	// Output: msisdn is invalid
}

func ExampleIsValidMsisdn() {
	fmt.Println(IsValidMsisdn("9955053777"))
	// Output: true
}

func ExampleIsValidMsisdn_fail() {
	fmt.Println(IsValidMsisdn("882123"))
	// Output: false
}
