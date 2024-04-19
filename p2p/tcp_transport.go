package p2p

import (
	"fmt"
	"net"
	"sync"
)

//INFO: represents a remote node in an established tcp connections
type TCPPeer struct {
	conn     net.Conn //INFO: underlying connection of the peer
	outbound bool     //INFO: basically lets us know if it was dialed or it's the dialer
}

type TCPTransportOps struct {
	ListenAddr    string
	HandShakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOps
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn,
		outbound,
	}
}

func NewTCPTransport(opts TCPTransportOps) *TCPTransport {
	return &TCPTransport{
		TCPTransportOps: opts,
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true) //INFO: create a peer that acccepts the connection

	if err := t.HandShakeFunc(peer); err != nil {
		fmt.Printf("TCP handshake error: %s\n", err)
		conn.Close()
		return
	}

	msg := &Temp{}

	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}

	}

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept loop error: %s\n", err) //TODO: pls better error handling
		}

		fmt.Printf("new incoming connection %+v\n", conn)

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}
