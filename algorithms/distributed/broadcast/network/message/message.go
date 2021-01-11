package message

func New(senderID int, msg string) *Message {
	return &Message{
		SenderID: senderID,
		Msg:      msg,
	}
}

type Message struct {
	SenderID int
	Msg      string
}
