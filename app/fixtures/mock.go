package fixtures

import "indexPrice/app/ticker"

type StorerMock struct {
	countClean int
	countMerge int
}

func (m *StorerMock) Clean() {
	m.countClean++
}

func (m *StorerMock) Merge() <-chan ticker.TickerPrice {
	m.countMerge++
	return AsChan("0", "2", "4")
}

func (m StorerMock) VerifyCleanInvoked() int {
	return m.countClean
}

func (m StorerMock) VerifyMergeInvoked() int {
	return m.countMerge
}
