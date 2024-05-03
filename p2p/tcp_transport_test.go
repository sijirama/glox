package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {

    opts := TCPTransportOps{
        ListenAddr: ":3000",
        HandShakeFunc: NOPHandshakeFunc,
        Decoder: DefaultDecoder{},
    }

    tr := NewTCPTransport(opts)

    assert.Equal(t, tr.ListenAddr, ":3000", "addresses should be equal bro")

	//NOTE: what does a transport always do, a transport always listens and accepts

	err := tr.ListenAndAccept()
	assert.Nil(t, err, err) // error should be nil
}
