package p2p

// make like a library
// Peer represetnts an interface the remote node
// representation of remote node

type Peer interface {
}

// Transport is anything that can handles the communication
// between nodes in the network
// This can be of the form (TCP, UDP, websockets...)
type Transport interface {
}
