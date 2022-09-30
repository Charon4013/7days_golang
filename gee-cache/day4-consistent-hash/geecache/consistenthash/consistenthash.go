package consistenthash

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// Map contains all the hashed keys
type Map struct {
	hash     Hash
	replicas int
	keys     []int //sorted
	hashMap  map[int]string
}
