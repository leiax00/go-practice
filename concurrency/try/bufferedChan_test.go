package try

import (
	"context"
	"golang.org/x/sync/errgroup"
	"sync"
	"testing"
)

func BenchmarkSize1(b *testing.B) {
	bufferedChan(b, 1)
}

func BenchmarkSize5(b *testing.B) {
	bufferedChan(b, 5)
}

func BenchmarkSize10(b *testing.B) {
	bufferedChan(b, 10)
}

func BenchmarkSize20(b *testing.B) {
	bufferedChan(b, 20)
}

func BenchmarkSize50(b *testing.B) {
	bufferedChan(b, 50)
}

func BenchmarkSize100(b *testing.B) {
	bufferedChan(b, 100)
}

func bufferedChan(b *testing.B, chanSize int) {
	var ch = make(chan int, chanSize)
	g, _ := errgroup.WithContext(context.Background())
	var sum = 0
	var lock sync.Mutex
	// producer
	g.Go(func() error {
		for i := 1; i <= 1000000; i++ {
			ch <- i
		}
		close(ch)
		return nil
	})

	// consumer
	for i := 0; i < 5; i++ {
		g.Go(func() error {
			var subSum = 0
			for item := range ch {
				subSum += item
			}
			lock.Lock()
			sum += subSum
			lock.Unlock()
			return nil
		})
	}
	err := g.Wait()
	if err != nil {
		return
	}
	//b.Log(sum)
}
