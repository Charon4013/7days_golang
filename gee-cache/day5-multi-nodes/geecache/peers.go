package geecache

// PeerPicker is an interface to locate the peer that owns a specific key
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is an interface that must be implemented by a peer
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
