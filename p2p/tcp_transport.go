package p2p

// why
import (
	"net"
	"sync"
)

type TCPTransport struct {
	//
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

/*func Tset() {
	// if we want to test this
	t := NewTCPTransport(":4344").(*TCPTransport)
	t.listener.Accept()
}*/
