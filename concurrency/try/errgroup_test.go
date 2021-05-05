package try

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
)

func TestErrGroup(t *testing.T) {
	errGroupMain()
}

func errGroupMain() {
	g, ctx := errgroup.WithContext(context.Background())
	var a, b, c []int
	g.Go(func() error {
		a = []int{1, 2, 3}
		//return errors.New("test")
		return nil
	})

	g.Go(func() error {
		b = []int{2, 3, 4}
		return nil
	})

	g.Go(func() error {
		c = []int{4, 5, 6}
		return nil
	})

	err := g.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ctx.Err())
	var sum = 0
	for i := 0; i < len(a); i++ {
		sum += a[i] + b[i] + c[i]
	}
	fmt.Println(sum)
}
