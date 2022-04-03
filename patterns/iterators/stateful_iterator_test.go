package iterators

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestStatefulIterator(t *testing.T) {
	t.Parallel()

	var arrays = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{101},
		{100, 99, 98, 97, 96, 95},
		{0, -1, -2, -3},
		{},
		{10000000},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	si := StatefulIteratorFromArrays(arrays...)

	for v := range si.Cycle(ctx) {
		<-time.After(500 * time.Millisecond)
		fmt.Println(v, si.HasNext())
	}

	fmt.Println("===================================================================")
	fmt.Println(si.last, si.exhausted)

	for i := 0; i < 10; i++ {
		v, ok := si.Next()
		fmt.Println(v, ok)
	}
}
