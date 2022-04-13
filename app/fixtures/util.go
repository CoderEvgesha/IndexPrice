package fixtures

import (
	"indexPrice/app/ticker"
	"time"
)

func AsChan(strs ...string) <-chan ticker.TickerPrice {
	c := make(chan ticker.TickerPrice)
	go func() {
		for _, str := range strs {
			c <- ticker.TickerPrice{
				Ticker: ticker.BTCUSDTicker,
				Time:   time.Now(),
				Price:  str,
			}
		}
		close(c)
	}()
	return c
}
