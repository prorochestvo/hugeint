package storage

import "sync"

// Storage is a sequence collection of nodes
type Storage struct {
	lock sync.RWMutex
	size uint64
	head *Node
}

// Begin returns first node in the collection
func (s *Storage) Begin() *Node {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.head
}

// End returns node in the collection with offset from last.
func (s *Storage) End(offset uint64) *Node {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if offset > s.size {
		return nil
	}

	limit := s.size - offset
	result := s.head

	for n := s.head; n != nil && limit > 0; n, limit = n.Next, limit-1 {
		result = n
	}

	return result
}

// Size returns total count the nodes in the collection
func (s *Storage) Size() uint64 {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.size
}

// Pop removes last node from of collections
func (s *Storage) Pop() *Node {
	var node = s.End(1)
	if node == nil {
		return nil
	}

  s.lock.Lock()
  defer s.lock.Unlock()

	var result *Node = nil
	// erase last pointer
	if node.Next != nil {
		result = node.Next
		node.Next = nil
		s.size--
	} else
	// erase head pointer
	if node == s.head {
		result = node
		s.head = nil
		s.size = 0
	}

	return result
}

// Push appends new node in the collection into last position
func (s *Storage) Push(n uint32) *Node {
	result := &Node{Data: n}

	node := s.End(0)

  s.lock.Lock()
  defer s.lock.Unlock()

	// make new head pointer
	if node == nil {
		s.head = result
	} else
	// grow new tail pointer
	{
		node.Next = result
	}
	s.size++

	return result
}

// ForEach returns each node from the collection
func (s *Storage) ForEach(f func(index uint64, node *Node) bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	item := s.head
	for index := uint64(0); item != nil; index, item = index+1, item.Next {
		if f(index, item) != true {
			return
		}
	}
}

// Free discards any nodes from collections and resets counter
func (s *Storage) Free() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.head = nil
	s.size = 0
}
