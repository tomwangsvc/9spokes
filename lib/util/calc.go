package util

import (
	"9spokes/lib/dto"
)

func CalculateByAccountCategoryAndAccountTypesAndValueType(accountsData []dto.AccountData, accountCategory string, accountTypes []string, valueType string) float64 {
	var result float64
	for _, accountData := range accountsData {
		if (accountCategory == "" || accountData.AccountCategory == accountCategory) &&
			(len(accountTypes) == 0 || contains(accountTypes, accountData.AccountType) &&
				(valueType == "" || accountData.ValueType == valueType)) {
			result += accountData.TotalValue
		}
	}

	return result
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
