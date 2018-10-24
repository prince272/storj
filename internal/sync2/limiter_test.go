// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information

package sync2

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestLimiterLimiting(t *testing.T) {
	const N, Limit = 1000, 10
	ctx := context.Background()
	limiter := NewLimiter(Limit)
	counter := int32(0)
	for i := 0; i < N; i++ {
		limiter.Go(ctx, func() {
			if atomic.AddInt32(&counter, 1) > Limit {
				panic("limit exceeded")
			}
			time.Sleep(time.Millisecond)
			atomic.AddInt32(&counter, -1)
		})
	}
	limiter.Wait()
}

func TestLimiterCancelling(t *testing.T) {
	const N, Limit = 1000, 10
	limiter := NewLimiter(Limit)

	ctx, cancel := context.WithCancel(context.Background())

	counter := int32(0)

	waitForCancel := make(chan struct{}, N)
	block := make(chan struct{})
	allreturned := make(chan struct{})

	go func() {
		for i := 0; i < N; i++ {
			limiter.Go(ctx, func() {
				if atomic.AddInt32(&counter, 1) > Limit {
					panic("limit exceeded")
				}

				waitForCancel <- struct{}{}
				<-block
			})
		}
		close(allreturned)
	}()

	for i := 0; i < Limit; i++ {
		<-waitForCancel
	}
	cancel()
	<-allreturned
	close(block)

	limiter.Wait()
	if counter > Limit {
		t.Fatal("too many times run")
	}
}
