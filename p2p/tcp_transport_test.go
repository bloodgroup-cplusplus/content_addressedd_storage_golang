package p2p

import "testing"

func TestTCPTransport(t *testing.T) {

	listenAddr := "4000"
	lr := NewTCPTransport(listenAddr)

}
