package util

import (
	"fmt"
	"math"
)

/*
Currency is prefixed with a $ sign
A comma is used to separate every 3 digits in the thousands, millions, billions, and trillions
Cents are removed
*/
func FormatCurrency(currency float64) string {
	currencyString := fmt.Sprintf("%.0f", (math.Trunc(currency))) // Remove cents
	newCurrencyString := currencyString
	i := 2
	ii := 2
	for ; i < len(currencyString); i++ {
		if i%3 == 0 {
			newCurrencyString = fmt.Sprintf("%s,%s", newCurrencyString[:(len(newCurrencyString)-ii)], newCurrencyString[(len(newCurrencyString)-ii):])
			ii++
		}
		ii++
	}
	return fmt.Sprintf("$%s", newCurrencyString) // Currency is prefixed with a $ sign
}

/*
All percentage values must be formatted to one decimal digit and be prefixed with a % sign.
*/
func FmtPercentage(c float64) string {
	return fmt.Sprintf("%.1f%s", math.Floor(c*10000)/100, "%")
}
