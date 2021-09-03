package util

import (
	"testing"
)

func Test_FormatCurrency(t *testing.T) {
	var data = []struct {
		Input  float64
		Expect string
	}{
		{
			Input:  1234,
			Expect: "$1,234",
		},
		{
			Input:  1234.23,
			Expect: "$1,234",
		},
		{
			Input:  12345678.23,
			Expect: "$12,345,678",
		},
		{
			Input:  12345678.090823,
			Expect: "$12,345,678",
		},
		{
			Input:  21412412345678.0909023,
			Expect: "$21,412,412,345,678",
		},
	}

	for i, d := range data {
		if result := FormatCurrency(d.Input); result != d.Expect {
			t.Errorf("Test failed in number %d, expected=%s, but result=%s", i, d.Expect, result)
		}
	}
}

func Test_FmtPercentage(t *testing.T) {
	var data = []struct {
		Input  float64
		Expect string
	}{
		{
			Input:  0.12,
			Expect: "12.0%",
		},
		{
			Input:  0.123,
			Expect: "12.3%",
		},
		{
			Input:  0.128,
			Expect: "12.8%",
		},
		{
			Input:  0.1287,
			Expect: "12.9%",
		},
	}

	for i, d := range data {
		if result := FmtPercentage(d.Input); result != d.Expect {
			t.Errorf("Test failed in number %d, expected=%s, but result=%s", i, d.Expect, result)
		}
	}
}
