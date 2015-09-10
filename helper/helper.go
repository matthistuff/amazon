package helper

import (
	"regexp"
)

// ASINReg matches a ASIN
var ASINReg, _ = regexp.Compile("[A-Za-z0-9]{10}")

// IsASIN checks wether str is an ASIN
func IsASIN(str string) bool {
	return ASINReg.MatchString(str)
}

// CartItemIDReg matches a cart item ID
var CartItemIDReg, _ = regexp.Compile("[A-Za-z0-9]{14}")

// IsCartItemID checks wether str is a cart item ID
func IsCartItemID(str string) bool {
	return CartItemIDReg.MatchString(str)
}

// NumericReg matches a numeric string
var NumericReg, _ = regexp.Compile("[0-9]+")

// IsNumeric checks wether str is numeric
func IsNumeric(str string) bool {
	return NumericReg.MatchString(str)
}
