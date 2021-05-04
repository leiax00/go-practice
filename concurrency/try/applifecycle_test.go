package try

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch:   make(chan string, 10),
		stop: make(chan struct{}),
	}
}

func (t *Tracker) Run() {
	for item := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(item)
	}
	//time.Sleep(5 * time.Second)  // make shutdown timeout
	t.stop <- struct{}{}
}

func (t *Tracker) Commit(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
		fmt.Println("stop")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}

func TestAppCycleDemo(t *testing.T) {
	tracker := NewTracker()
	go tracker.Run()
	_ = tracker.Commit(context.Background(), "tracker_01")
	_ = tracker.Commit(context.Background(), "tracker_02")
	_ = tracker.Commit(context.Background(), "tracker_03")
	time.Sleep(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	tracker.Shutdown(ctx)

}
