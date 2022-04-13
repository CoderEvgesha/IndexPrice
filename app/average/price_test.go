package average

import (
	"github.com/shopspring/decimal"
	"indexPrice/app/fixtures"
	"math/rand"
	"sync"
	"testing"
)

func TestCalculate(t *testing.T) {

	type test struct {
		input []string
		want  decimal.Decimal
	}

	tests := []test{
		{input: []string{"0", "10", "12.2", "13.2345122"}, want: decimal.NewFromFloat(6.929314025)},
		{input: []string{"-1", "-10", "-12.2", "-13.2345122"}, want: decimal.NewFromFloat(-1.0896570125)},
		{input: []string{"0"}, want: decimal.NewFromFloat(-0.54482850625)},
		{input: []string{"abc", ".ddkdu2", "234js"}, want: decimal.NewFromFloat(-0.272414253125)},
		{input: []string{"abc", "-2", "4.0"}, want: decimal.NewFromFloat(0.3637928734375)},
	}

	price := Price{sync.RWMutex{}, decimal.NewFromInt(5)}

	for _, ts := range tests {
		price.Calculate(fixtures.AsChan(ts.input...))
		if !price.price.Equal(ts.want) {
			t.Errorf("error for %v, result should be %v, but %v", ts.input, ts.want, price.price)
		}
	}
}

func TestGet(t *testing.T) {
	price := NewAveragePrice()

	for i := 0; i < 10; i++ {
		test := decimal.NewFromInt(int64(rand.Intn(100)))
		price.price = test
		if !price.Get().Equal(test) {
			t.Errorf("result should be %v", test)
		}
	}
}
