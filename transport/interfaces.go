package transport

import "udp_sim/core"

type Sender interface {
	Send(msg core.PoseMessage) error
}

type Receiver interface {
	Start()
}

type Encoder interface {
	Encode(msg core.PoseMessage) ([]byte, error)
	Decode(data []byte) (core.PoseMessage, error)
}

type MessageHandler interface {
	Handle(msg core.PoseMessage)
}
