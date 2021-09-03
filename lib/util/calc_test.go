package util

import (
	"9spokes/lib/dto"
	"testing"
)

func Test_contain(t *testing.T) {
	type input struct {
		s  []string
		ss string
	}
	var data = []struct {
		Desc   string
		Input  input
		Expect bool
	}{
		{
			Desc: "empty",
			Input: input{
				s:  nil,
				ss: "",
			},
			Expect: false,
		},
		{
			Desc: "empty but has value",
			Input: input{
				s:  nil,
				ss: "s",
			},
			Expect: false,
		},
		{
			Desc: "contains",
			Input: input{
				s:  []string{"s"},
				ss: "s",
			},
			Expect: true,
		},
		{
			Desc: "contains by more values",
			Input: input{
				s:  []string{"s", "s"},
				ss: "s",
			},
			Expect: true,
		},
	}

	for i, d := range data {
		if result := contains(d.Input.s, d.Input.ss); result != d.Expect {
			t.Errorf("Test failed in number %d, expected=%t, but result=%t", i, d.Expect, result)
		}
	}
}

func Test_CalculateByAccountCategoryAndAccountTypesAndValueType(t *testing.T) {
	type Input struct {
		AccountsData    []dto.AccountData
		AccountCategory string
		AccountTypes    []string
		ValueType       string
	}

	var data = []struct {
		Desc   string
		Input  Input
		Expect float64
	}{
		{
			Desc: "default",
			Input: Input{
				AccountsData:    []dto.AccountData{},
				AccountCategory: "",
				AccountTypes:    nil,
				ValueType:       "",
			},
			Expect: 0,
		},
		{
			Desc: "contain by AccountCategory",
			Input: Input{
				AccountsData: []dto.AccountData{
					{
						AccountCategory: "AccountCategory",
						TotalValue:      1,
					},
				},
				AccountCategory: "AccountCategory",
				AccountTypes:    nil,
				ValueType:       "",
			},
			Expect: 1,
		},
		{
			Desc: "contain by AccountType",
			Input: Input{
				AccountsData: []dto.AccountData{
					{
						AccountType: "AccountType",
						TotalValue:  1,
					},
				},
				AccountCategory: "",
				AccountTypes:    []string{"AccountType"},
				ValueType:       "",
			},
			Expect: 1,
		},
		{
			Desc: "contain by ValueType",
			Input: Input{
				AccountsData: []dto.AccountData{
					{
						TotalValue: 1,
						ValueType:  "ValueType",
					},
				},
				AccountCategory: "",
				AccountTypes:    []string{""},
				ValueType:       "ValueType",
			},
			Expect: 1,
		},
		{
			Desc: "one contains by AccountCategory and ValueType, one not match",
			Input: Input{
				AccountsData: []dto.AccountData{
					{
						AccountCategory: "AccountCategory",
						TotalValue:      2,
						ValueType:       "ValueType",
					},
					{
						TotalValue: 1,
						ValueType:  "ValueType",
					},
				},
				AccountCategory: "AccountCategory",
				AccountTypes:    []string{""},
				ValueType:       "ValueType",
			},
			Expect: 2,
		},
		{
			Desc: "one contains by AccountCategory and AccountType, one not match",
			Input: Input{
				AccountsData: []dto.AccountData{
					{
						AccountCategory: "AccountCategory",
						AccountType:     "AccountType",
						TotalValue:      2,
						ValueType:       "ValueType",
					},
					{
						TotalValue: 1,
						ValueType:  "ValueType",
					},
				},
				AccountCategory: "AccountCategory",
				AccountTypes:    []string{"AccountType"},
				ValueType:       "ValueType",
			},
			Expect: 2,
		},
		{
			Desc: "one contains by AccountCategory, AccountTypes and ValueType, and one not match",
			Input: Input{
				AccountsData: []dto.AccountData{
					{
						AccountCategory: "AccountCategory",
						AccountType:     "AccountType",
						TotalValue:      2,
						ValueType:       "ValueType",
					},
					{
						TotalValue: 1,
						ValueType:  "ValueType",
					},
				},
				AccountCategory: "AccountCategory",
				AccountTypes:    []string{"AccountType"},
				ValueType:       "ValueType",
			},
			Expect: 2,
		},
	}

	for i, d := range data {
		if result := CalculateByAccountCategoryAndAccountTypesAndValueType(d.Input.AccountsData, d.Input.AccountCategory, d.Input.AccountTypes, d.Input.ValueType); result != d.Expect {
			t.Errorf("Test failed in number %d, expected=%f, but result=%f", i, d.Expect, result)
		}
	}
}
