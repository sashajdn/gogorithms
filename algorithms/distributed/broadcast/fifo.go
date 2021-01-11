package broadcast

import (
	"fmt"
	"gogorithms/dalgorithms/broadcast/network"
	"gogorithms/dalgorithms/broadcast/network/message"
	"gogorithms/structures/clocks"
	"gogorithms/structures/queues"
	"sync"
)

type FIFO struct {
	// Id is the Node ID that will implement the FIFO broadcast algo
	Id int
	// SendSeq is the monotonic send sequence of messages
	SendSeq int

	mu  sync.Locker
	c   *clocks.VectorClock
	buf *Buffer
	n   network.Network
}

// New is a factory to create a new FIFO Impl.
func New(identifier int, numberOfNodes int, maxBufSize int) *FIFO {
	return &FIFO{
		Id:      identifier,
		SendSeq: 0,
		c:       clocks.New(identifier, numberOfNodes),
		buf: &Buffer{
			q: queues.NewBasicQueue(maxBufSize),
		},
		n: network.New(),
	}
}

// Broadcast broadcasts a message via the interval network.
func (f *FIFO) Broadcast(msg *message.Message) error {
	err := f.n.Broadcast(f.Id, f.SendSeq, msg)
	if err != nil {
		return err
	}
	f.incrementSendSeq()
	return nil
}

// StartReceiver starts a new receiver server to receive new messages from.
func (f *FIFO) StartReceiver() <-chan error {
	errCh := make(chan error)
	go func() {
		for msg := range f.n.Receiver() {
			err := f.receive(msg)
			if err != nil {
				errCh <- err
			}
		}
	}()
	return errCh
}

func (f *FIFO) incrementSendSeq() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.SendSeq++
}

func (f *FIFO) receive(msg *message.Message) error {
	ok := f.buf.q.Enqueue(msg)
	if !ok {
		return fmt.Errorf("Failed to append to internal buffer of size %v", f.buf.q.Size())
	}
	for msg := range f.buf.getWaitingMessages() {
		msg, ok := msg.(message.Message)
		if !ok {
			return fmt.Errorf("Expecting a message of type `message.Message` but failed.")
		}
		f.c.Increment(msg.SenderID)
	}
	return nil
}
