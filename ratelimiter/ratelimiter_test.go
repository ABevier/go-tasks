package ratelimiter

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRateLimiter(t *testing.T) {
	require := require.New(t)

	wg := sync.WaitGroup{}

	run := func(ctx context.Context, n int) (int, error) {
		return n * 2, nil
	}

	rl := New(RateLimiterOpts{Limit: 10, Burst: 1, FullQueueStrategy: BlockWhenFull}, run)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			r, err := rl.Submit(context.Background(), n)
			require.NoError(err)
			require.Equal(n*2, r)
		}(i)
	}

	wg.Wait()
}
