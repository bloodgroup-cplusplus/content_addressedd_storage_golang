package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {

	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)

	assert.Equal(t, tr.listenAddress, listenAddr)

	//Server
	// tr.start()

	// the transporter always listens and accepts

	assert.Nil(t, tr.ListenAndAccept())

}
