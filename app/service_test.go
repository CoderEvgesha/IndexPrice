package app

import (
	"context"
	"indexPrice/app/fixtures"
	"math"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	mock := fixtures.StorerMock{}
	service := NewIndexService(&mock)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	service.Run(ctx)

	clean := mock.VerifyCleanInvoked()
	merge := mock.VerifyMergeInvoked()

	if clean == 0 {
		t.Errorf("the сlean function should have been called")
	}

	if merge == 0 {
		t.Errorf("the merge function should have been called")
	}

	if math.Abs(float64(merge/100-clean)) > 1 {
		t.Errorf("wrong functions call frequency")
	}
}
