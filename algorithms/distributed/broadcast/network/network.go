package network

import "gogorithms/dalgorithms/broadcast/network/message"

type Network interface {
	Broadcast(id int, sendSeq int, msg *message.Message) error
	Receiver() <-chan *message.Message
}

// NetworkImpl is a synchronous network implemention
type NetworkImpl struct{}

func New() *NetworkImpl {
	return &NetworkImpl{}
}

func (n *NetworkImpl) Broadcast(id int, sendSeq int, msg *message.Message) error {
	return nil
}

func (n *NetworkImpl) Receiver() <-chan *message.Message {
	c := make(chan *message.Message)
	return c
}
