package ticker

import "time"

const (
	BTCUSDTicker Ticker = "BTC_USD"
)

type Ticker string

type TickerPrice struct {
	Ticker Ticker
	Time   time.Time
	Price  string
}
