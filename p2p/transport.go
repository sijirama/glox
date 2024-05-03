package p2p

//INFO: Peer is an interface that represents a remote node
type Peer interface {
	Close() error
}

//INFO: Transport handles all communication between nodes in network
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
