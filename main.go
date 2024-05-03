package main

import (
	"fmt"
	"log"

	"github.com/sijirama/glox/p2p"
)

func OnPeer(peer p2p.Peer) error {
	fmt.Println("Processing some logic with the peer outside the TCP Transport")
	//peer.Close()
	return nil
}

func main() {

	opts := p2p.TCPTransportOps{
		ListenAddr:    ":3000",
		HandShakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTCPTransport(opts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
