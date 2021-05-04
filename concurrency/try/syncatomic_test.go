package try

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkOriginFn(b *testing.B) {
	OriginFn()
}

func BenchmarkAtomic(b *testing.B) {
	AtomicFn()
}

func BenchmarkMutex(b *testing.B) {
	MutexFn()
}

func BenchmarkRWMutex(b *testing.B) {
	RWMutexFn()
}

type Config struct {
	data []int
}

// OriginFn 存在竞争的方式
func OriginFn() {
	cfg := &Config{}

	go func() {
		i := 0
		for {
			i++
			// write val, not atomic
			cfg.data = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				// read val
				fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// AtomicFn 使用 atomic.Value 保证原子性
func AtomicFn() {
	var v atomic.Value
	v.Store(&Config{})

	go func() {
		i := 0
		for {
			i++
			// write val, atomic
			cfg := &Config{[]int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				// read val
				cfg := v.Load().(*Config)
				fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// MutexFn 使用 sync.Mutex 保证原子性
func MutexFn() {
	var lock sync.Mutex
	cfg := &Config{}

	go func() {
		i := 0
		for {
			i++
			lock.Lock()
			// write val, not atomic
			cfg.data = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
			lock.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				// read val
				lock.Lock()
				fmt.Printf("%v\n", cfg)
				lock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// RWMutexFn 使用 sync.RWMutex 保证原子性
func RWMutexFn() {
	var lock sync.RWMutex
	cfg := &Config{}

	go func() {
		i := 0
		for {
			i++
			lock.Lock()
			// write val, not atomic
			cfg.data = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
			lock.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				// read val
				lock.RLock()
				fmt.Printf("%v\n", cfg)
				lock.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
