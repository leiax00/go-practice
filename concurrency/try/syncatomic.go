package try

import (
	"fmt"
	"sync"
)

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
