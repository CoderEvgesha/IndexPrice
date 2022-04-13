package sources

import "indexPrice/app/ticker"

type PriceStreamSubscriber interface {
	SubscribePriceStream(ticker.Ticker) (chan ticker.TickerPrice, chan error)
}
