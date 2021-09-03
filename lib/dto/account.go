package dto

import "time"

type AccountRawData struct {
	BalanceDate          time.Time     `json:"balance_date"`
	ConnectionId         string        `json:"connection_id"`
	Data                 []AccountData `json:"data"`
	Currency             string        `json:"currency"`
	ObjectCategory       string        `json:"object_category"`
	ObjectClass          string        `json:"object_class"`
	ObjectCreationDate   string        `json:"object_creation_date"`
	ObjectOriginCategory string        `json:"object_origin_category"`
	ObjectOriginType     string        `json:"object_origin_type"`
	ObjectType           string        `json:"object_type"`
	User                 string        `json:"user"`
}

type AccountData struct {
	AccountCategory   string  `json:"account_category"`
	AccountCode       string  `json:"account_code"`
	AccountCurrency   string  `json:"account_currency"`
	AccountIdentifier string  `json:"account_identifier"`
	AccountName       string  `json:"account_name"`
	AccountStatus     string  `json:"account_status"`
	AccountType       string  `json:"account_type"`
	AccountTypeBank   string  `json:"account_type_bank"`
	SystemAccount     string  `json:"system_account"`
	TotalValue        float64 `json:"total_value"`
	ValueType         string  `json:"value_type"`
}

type Result struct {
	/*
		This should be calculated by adding up all the values under total_value where the account_category field is set to revenue
	*/
	Revenue string `json:"revenue"`
	/*
		This should be calculated by adding up all the values under total_value where the account_category field is set to expense
	*/
	Expenses string `json:"expenses"`
	/*
		This is calculated in two steps: first by adding all the total_value fields where the account_type is set to sales and the value_type is set to debit; then dividing that by the revenue value calculated earlier to generate a percentage value.
	*/
	GrossProfitMargin string `json:"gross_profit_margin"`
	/*
		This metric is calculated by subtracting the expenses value from the revenue value and dividing the remainder by revenue to calculate a percentage.
	*/
	NetProfitMargin string `json:"net_profit_margin"`
	/*
		This is calculated dividing the assets by the liabilities creating a percentage value where assets are calculated by:

		adding the total_value from all records where the account_category is set to assets, the value_type is set to debit, and the account_type is one of current, bank, or current_accounts_receivable
		subtracting the total_value from all records where the account_category is set to assets, the value_type is set to credit, and the account_type is one of current, bank, or current_accounts_receivable
		and liabilities are calculated by:

		adding the total_value from all records where the account_category is set to liability, the value_type is set to credit, and the account_type is one of current or current_accounts_payable
		subtracting the total_value from all records where the account_category is set to liability, the value_type is set to debit, and the account_type is one current or current_accounts_payable
	*/
	WorkingCapitalRatio string `json:"Working_capital_ratio"`
}
