package p2p

import "net"

//INFO: Message holds data being sent between nodes in the network
type Message struct {
    From net.Addr
	Payload []byte
}
