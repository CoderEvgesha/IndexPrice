package average

import (
	"github.com/shopspring/decimal"
	"indexPrice/app/ticker"
)

func calculate(ch <-chan ticker.TickerPrice) decimal.Decimal {
	count := int64(0)
	sum := decimal.NewFromFloat(0)
	for v := range ch {
		if fee, err := decimal.NewFromString(v.Price); err == nil {
			sum = sum.Add(fee)
			count++
		}
	}
	if sum.IsZero() {
		return sum
	}
	return sum.Div(decimal.NewFromInt(count))
}
