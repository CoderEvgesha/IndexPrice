package app

import (
	"context"
	"github.com/shopspring/decimal"
	"indexPrice/app/average"
	"indexPrice/app/ticker"
)

type Storer interface {
	Clean()
	Merge() <-chan ticker.TickerPrice
}

type IndexService struct {
	pool  Storer
	value *average.Price
}

func NewIndexService(pool Storer) *IndexService {
	return &IndexService{pool, average.NewAveragePrice()}
}

func (s *IndexService) Run(ctx context.Context) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			s.value.Calculate(s.pool.Merge())
			if i%100 == 0 {
				s.pool.Clean()
			}
		}
	}
}

func (s IndexService) GetFairPrice() decimal.Decimal {
	return s.value.Get()
}
