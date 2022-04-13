package average

import (
	"github.com/shopspring/decimal"
	"indexPrice/app/ticker"
	"sync"
)

type Price struct {
	mu    sync.RWMutex
	price decimal.Decimal
}

func NewAveragePrice() *Price {
	return &Price{sync.RWMutex{}, decimal.NewFromInt(0)}
}

func (p *Price) Get() decimal.Decimal {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.price
}

func (p *Price) Calculate(prices <-chan ticker.TickerPrice) {
	last := calculate(prices)
	average := last.Add(p.price).Div(decimal.NewFromInt(2))
	p.mu.Lock()
	p.price = average
	defer p.mu.Unlock()
}
