package storage

// Node is element of collection in the storage
type Node struct {
	Data uint32
	Next *Node
}
