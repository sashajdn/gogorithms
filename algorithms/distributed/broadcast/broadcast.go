package broadcast

import "gogorithms/algorithms/distributed/broadcast/network/message"

// Broadcaster interface that defines the broadcast algorithm behaviour.
type Broadcaster interface {
	Broadcast(msg *message.Message) error
	StartReceiver() <-chan error
}
