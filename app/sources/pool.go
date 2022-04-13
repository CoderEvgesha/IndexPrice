package sources

import (
	"context"
	"indexPrice/app/ticker"
	"sync"
	"time"
)

type Pool struct {
	mu   sync.RWMutex
	subs []*pair
}

type pair struct {
	prices   <-chan ticker.TickerPrice
	errors   <-chan error
	isClosed bool
}

func NewPool() *Pool {
	ps := &Pool{}
	ps.subs = make([]*pair, 100)
	return ps
}

func (s *Pool) GetSubscribers() []*pair {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.subs
}

func (s *Pool) AddSubscribers(prices <-chan ticker.TickerPrice, errors <-chan error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.subs = append(s.subs, &pair{prices, errors, false})
}

func (s *Pool) Clean() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, p := range s.subs {
		if p.isClosed {
			s.subs[i] = s.subs[len(s.subs)-1]
			s.subs[len(s.subs)-1] = nil
			s.subs = s.subs[:len(s.subs)-1]
		}
	}
}

func (p *Pool) Merge() <-chan ticker.TickerPrice {
	out := make(chan ticker.TickerPrice)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var wg sync.WaitGroup
		wg.Add(len(p.subs))

		for _, p := range p.subs {
			go func(ctx context.Context, p *pair) {
				defer wg.Done()
				for {
					if p.isClosed {
						return
					}
					select {
					case v, ok := <-p.prices:
						if !ok {
							p.isClosed = true
						}
						if time.Now().Sub(v.Time) < time.Minute {
							out <- v
						}
					case err, ok := <-p.errors:
						if !ok || err != nil {
							p.isClosed = true
						}
					case <-ctx.Done():
						return
					}
				}
			}(ctx, p)
		}
		wg.Wait()
		close(out)
	}()
	return out
}
