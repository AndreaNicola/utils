package utils

import (
	"errors"
	"github.com/ttacon/libphonenumber"
)

func PhoneNumberParseAndFormat(phoneNumber string) (string, error) {

	num, err := libphonenumber.Parse(phoneNumber, "IT")
	if err != nil {
		return "", errors.New("this phone number is not valid")
	}

	isValid := libphonenumber.IsValidNumber(num)
	countryCode := num.CountryCode
	numType := libphonenumber.GetNumberType(num)

	if isValid && *countryCode == 39 && numType == libphonenumber.MOBILE {
		return libphonenumber.Format(num, libphonenumber.INTERNATIONAL), nil
	} else {
		return "", errors.New("this phone number is not italian nor a valid mobile number")
	}

}
