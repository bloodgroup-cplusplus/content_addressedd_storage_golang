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
