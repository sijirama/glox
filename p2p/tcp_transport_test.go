package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {

	listenAddr := ":4000"
	tr := NewTCPTransport(TCPTransportOps{
		ListenAddr: listenAddr,
	})
	assert.Equal(t, tr.ListenAddr, listenAddr, "addresses should be equal bro")

	//NOTE: what does a transport always do, a transport always listens and accepts

	err := tr.ListenAndAccept()
	assert.Nil(t, err, err) // error should be nil

	select {}
}
