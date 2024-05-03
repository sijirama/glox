package p2p

import "net"

//INFO: RPC holds data being sent between nodes in the network
type RPC struct {
    From net.Addr
	Payload []byte
}
