package generators

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestIterator(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	// iterator := make(Iterator, 1)
	// gen := iterator.Generator(ctx, 1, 10, 20)

	//for v := range gen() {
	//	fmt.Println(v)
	//	time.Sleep(200 * time.Millisecond)
	//}

	//for v := range Negative(ctx, gen) {
	//	fmt.Println(v)
	//	time.Sleep(200 * time.Millisecond)
	//}

	//for v := range Cycle(ctx, gen) {
	//	fmt.Println(v)
	//	time.Sleep(200 * time.Millisecond)
	//}

	iterator1 := make(Iterator, 1)

	iterator1 = iterator1.Range(ctx, 5, 0, 100)

	for v := range Sink(ctx, iterator1) {
		fmt.Println(v)
		time.Sleep(200 * time.Millisecond)
	}
}
