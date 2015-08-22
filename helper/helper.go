package helper

import (
	"regexp"
)

var ASINReg, _ = regexp.Compile("[A-Za-z0-9]{10}")

func IsASIN(str string) bool {
	return ASINReg.MatchString(str)
}

var CartItemIdReg, _ = regexp.Compile("[A-Za-z0-9]{14}")

func IsCartItemId(str string) bool {
	return CartItemIdReg.MatchString(str)
}

var NumericReg, _ = regexp.Compile("[0-9]+")

func IsNumeric(str string) bool {
	return NumericReg.MatchString(str)
}
