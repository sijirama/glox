package main

import (
	"github.com/sijirama/glox/p2p"
	"log"
)

func main() {

	opts := p2p.TCPTransportOps{
		ListenAddr:    ":3000",
		HandShakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	tr := p2p.NewTCPTransport(opts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
