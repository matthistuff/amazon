package helpers
import (
	"regexp"
)

var ASINReg, _ = regexp.Compile("[A-Za-z0-9]{10}")

func IsASIN(str string) bool {
	return ASINReg.MatchString(str)
}

var NumericReg, _ = regexp.Compile("[0-9]+")

func IsNumeric(str string) bool {
	return NumericReg.MatchString(str)
}