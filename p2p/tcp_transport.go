package p2p

import (
	"fmt"
	"net"
	"sync"
)

//INFO: represents a remote node in an established tcp connections
type TCPPeer struct {
	conn     net.Conn //INFO: underlying connection of the peer
	outbound bool     //INFO: basically lets us know if it was dialed ot it's the dialer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn,
		outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NOPHandshakeFunc,
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true) //INFO: create a peer that acccepts the connection

	if err := t.shakeHands(conn); err != nil {

	}

	fmt.Printf("new incoming connection %+v\n", peer)
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept loop error: %s\n", err) //TODO: pls better error handling
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}
