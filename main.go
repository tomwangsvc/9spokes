package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"9spokes/lib/dto"
	"9spokes/lib/errors"
	"9spokes/lib/util"
)

func main() {
	dataBytes, err := os.ReadFile("data.json")
	if err != nil {
		log.Panic("data json can't be read")
	}

	result, err := process(dataBytes)
	if err != nil {
		log.Panicf("Failed processing by %v", err.Error())
	}

	log.Println(fmt.Sprintf("\n Revenue: %s\n Expenses: %s\n Gross Profit Margin: %s\n Net Profit Margin: %s\n Working Capital Ratio: %s", result.Revenue, result.Expenses, result.GrossProfitMargin, result.NetProfitMargin, result.WorkingCapitalRatio))
}

func process(dataBytes []byte) (*dto.Result, error) {
	var accountRawData dto.AccountRawData
	if err := json.Unmarshal(dataBytes, &accountRawData); err != nil {
		return nil, errors.CustomError{Desc: "Failed unmarshalling by json", RawErr: err}
	}

	revenue := util.CalculateByAccountCategoryAndAccountTypesAndValueType(accountRawData.Data, "revenue", nil, "")
	expense := util.CalculateByAccountCategoryAndAccountTypesAndValueType(accountRawData.Data, "expense", nil, "")
	assets := util.CalculateByAccountCategoryAndAccountTypesAndValueType(accountRawData.Data, "assets", []string{"current", "bank", "current_accounts_receivable"}, "debit") - util.CalculateByAccountCategoryAndAccountTypesAndValueType(accountRawData.Data, "assets", []string{"current", "bank", "current_accounts_receivable"}, "credit")
	liabilities := util.CalculateByAccountCategoryAndAccountTypesAndValueType(accountRawData.Data, "liability", []string{"current", "current_accounts_payable"}, "credit") - util.CalculateByAccountCategoryAndAccountTypesAndValueType(accountRawData.Data, "liability", []string{"current", "current_accounts_payable"}, "debit")

	return &dto.Result{
		Revenue:  util.FormatCurrency(revenue),
		Expenses: util.FormatCurrency(expense),
		GrossProfitMargin: util.FmtPercentage(
			util.CalculateByAccountCategoryAndAccountTypesAndValueType(
				accountRawData.Data,
				"",
				[]string{"sales"},
				"debit",
			) / revenue),
		NetProfitMargin: util.FmtPercentage((revenue - expense) / revenue),
		WorkingCapitalRatio: util.FmtPercentage(
			assets / liabilities,
		),
	}, nil
}
