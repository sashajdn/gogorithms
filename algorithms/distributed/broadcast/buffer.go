package broadcast

import "gogorithms/structures/queues"

// Buffer is a wrapper around a queue. With the added internal behaviour of pumping
// messages into a channel.
type Buffer struct {
	// thread safe basic queue.
	q *queues.BasicQueue
}

func (b *Buffer) getWaitingMessages() chan interface{} {
	c := make(chan interface{})
	defer close(c)
	go func() {
		for {
			msg, ok := b.q.Pop()
			if !ok {
				break
			}
			c <- msg
		}
	}()
	return c
}
