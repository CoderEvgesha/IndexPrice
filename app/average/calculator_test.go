package average

import (
	"github.com/shopspring/decimal"
	"indexPrice/app/fixtures"
	"testing"
)

func TestCalculateAverage(t *testing.T) {

	type test struct {
		input []string
		want  decimal.Decimal
	}

	tests := []test{
		{input: []string{"0", "10", "12.2", "13.2345122"}, want: decimal.NewFromFloat(8.85862805)},
		{input: []string{"-1", "-10", "-12.2", "-13.2345122"}, want: decimal.NewFromFloat(-9.10862805)},
		{input: []string{"0"}, want: decimal.NewFromInt(0)},
		{input: []string{"abc", ".ddkdu2", "234js"}, want: decimal.NewFromInt(0)},
		{input: []string{"abc", "-2", "4.0"}, want: decimal.NewFromInt(1)},
	}

	for _, ts := range tests {
		if res := calculate(fixtures.AsChan(ts.input...)); !res.Equal(ts.want) {
			t.Errorf("error for %v, result should be %v, but %v", ts.input, ts.want, res)
		}
	}
}
