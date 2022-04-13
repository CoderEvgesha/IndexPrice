package sources

import (
	"indexPrice/app/fixtures"
	"sync"
	"testing"
)

func TestGetSubscribers(t *testing.T) {
	pairs := []*pair{
		{fixtures.AsChan("1", "4.6", "0", "_123", "10"), make(chan error), false},
	}

	pool := Pool{sync.RWMutex{}, pairs}

	res := pool.GetSubscribers()
	if len(res) != 1 && pairs[0] != res[0] {
		t.Errorf("getting subscribers does not work correctly")
	}
}

func TestAddSubscribers(t *testing.T) {
	pool := Pool{sync.RWMutex{}, []*pair{}}
	sub := pair{fixtures.AsChan("5", "5", "5"), make(chan error), false}

	pool.AddSubscribers(sub.prices, sub.errors)

	if len(pool.subs) == 0 {
		t.Errorf("add method didn't work")
	}

	if sub != *pool.subs[0] {
		t.Errorf("add method saved wrong entity")
	}
}

func TestClean(t *testing.T) {
	pairs := []*pair{
		{fixtures.AsChan("1", "4.6", "0", "_123", "10"), make(chan error), false},
		{fixtures.AsChan("5", "5", "5"), make(chan error), true},
	}

	pool := Pool{sync.RWMutex{}, pairs}
	pool.Clean()

	if len(pool.subs) > 1 {
		t.Errorf("clean method didn't work")
	}
}

func TestMerge(t *testing.T) {
	pairs := []*pair{
		{fixtures.AsChan("1"), make(chan error), false},
		{fixtures.AsChan("-5"), make(chan error), false},
		{fixtures.AsChan("_123"), make(chan error), true},
	}
	pool := Pool{sync.RWMutex{}, pairs}

	res := pool.Merge()

	i := 0
	for v := range res {
		if v.Price != "1" && v.Price != "-5" {
			t.Errorf("wrong values in the result channel")
		}
		i++
	}

	if i != 2 {
		t.Errorf("wrong number of elements in the result channel")
	}

}
