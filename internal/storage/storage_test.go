package storage

import (
	"fmt"
  "sync"
  "testing"
)

func TestStorage_Push(t *testing.T) {
	// init Storage: 1, 2, 3, 4, 5, 6, 7, 8
	s := Storage{ lock: sync.RWMutex{} }
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)

	// check nodes position
	var index uint32 = 0
	item := s.head
	for item != nil {
		index++
		if item.Data != index {
			t.Errorf("storage[%d] is not correctly; expected: %d; actual: %d;", index-1, index, item.Data)
		}
		item = item.Next
	}

	// check size
	if index != 8 {
		t.Errorf("length is not correctly; expected: %d; actual: %d;", 8, index)
	}
	if s.size != 8 {
		t.Errorf("length is not correctly; expected: %d; actual: %d;", 8, s.size)
	}

	// check head (first pointer)
	if s.head == nil {
		t.Errorf("first pointer of nodes is not correctly; expected: %s; actual: %s;", "not nil", "nil")
	}
}

func TestStorage_Pop(t *testing.T) {
	// init Storage: 1, 2, 3, 4, 5, 6, 7, 8
	s := Storage{ lock: sync.RWMutex{} }
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)

	// remove last 4 nodes
	if n := s.Pop(); n == nil || n.Data != 8 {
		d := "nil"
		if n != nil {
			d = fmt.Sprintf("%d", n.Data)
		}
		t.Errorf("last node data is not correctly; expected: %d; actual: %v;", 8, d)
	}
	if n := s.Pop(); n == nil || n.Data != 7 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("last node data is not correctly; expected: %d; actual: %v;", 7, d)
	}
	if n := s.Pop(); n == nil || n.Data != 6 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("last node data is not correctly; expected: %d; actual: %v;", 6, d)
	}
	if n := s.Pop(); n == nil || n.Data != 5 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("last node data is not correctly; expected: %d; actual: %v;", 5, d)
	}

	// check size
	var index uint32 = 0
	item := s.head
	for item != nil {
		item = item.Next
		index++
	}
	if index != 4 {
    t.Errorf("length is not correctly; expected: %d; actual: %d;", 4, index)
	}
	if s.size != 4 {
    t.Errorf("length is not correctly; expected: %d; actual: %d;", 4, s.size)
	}
	if s.head == nil {
		t.Errorf("first pointer of nodes is not correctly; expected: %s; actual: %s;", "not nil", "nil")
	}

	// remove last 4 nodes
	if n := s.Pop(); n == nil || n.Data != 4 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("last node data is not correctly; expected: %d; actual: %v;", 4, d)
	}
	if n := s.Pop(); n == nil || n.Data != 3 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("last node data is not correctly; expected: %d; actual: %v;", 3, d)
	}
	if n := s.Pop(); n == nil || n.Data != 2 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("last node data is not correctly; expected: %d; actual: %v;", 2, d)
	}
	if n := s.Pop(); n == nil || n.Data != 1 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("last node data is not correctly; expected: %d; actual: %v;", 1, d)
	}

	// check size
	if s.size != 0 {
		t.Errorf("length is not correctly; expected: %d; actual: %d;", 0, s.size)
	}
	if s.head != nil {
		t.Errorf("first pointer of nodes is not correctly; expected: %s; actual: %s;", "nil", "not nil")
	}
}

func TestStorage_End(t *testing.T) {
	// init storage: 1, 2, 3, 4, 5, 6, 7, 8
	s := Storage{ lock: sync.RWMutex{} }
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)

	// check tail
	n := s.End(0)
	if n == nil || n.Data != 8 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("node is wrong; expected: %d; actual: %s;", 8, d)
	}

	n = s.End(1)
	if n == nil || n.Data != 7 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("node is wrong; expected: %d; actual: %s;", 7, d)
	}

	n = s.End(2)
	if n == nil || n.Data != 6 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("node is wrong; expected: %d; actual: %s;", 6, d)
	}

	n = s.End(4)
	if n == nil || n.Data != 4 {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("node is wrong; expected: %d; actual: %s;", 4, d)
	}

	n = s.End(8)
	if n == nil || n.Data != 1 || n != s.head {
    d := "nil"
    if n != nil {
      d = fmt.Sprintf("%d", n.Data)
    }
    t.Errorf("node is wrong; expected: %d; actual: %s;", 1, d)
	}

	n = s.End(10)
	if n != nil {
    t.Errorf("node is wrong; expected: %s; actual: %d;", "nil", n.Data)
	}

	if s.size != 8 {
    t.Errorf("length is not correctly; expected: %d; actual: %d;", 0, s.size)
	}
	if s.head == nil {
    t.Errorf("first pointer of nodes is not correctly; expected: %s; actual: %s;", "not nil", "nil")
	}
}

func TestStorage_ForEach(t *testing.T) {
	index := 0
	s := Storage{ lock: sync.RWMutex{} }
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)
	s.Push(9)
	s.Push(10)

	s.ForEach(func(i uint64, n *Node) bool {
		if d := uint64(n.Data); i != d {
      t.Errorf("storage[%d] is not correctly; expected: %d; actual: %d;", index, i, d)
		}
		index++
		return true
	})

	if index != 11 {
    t.Errorf("length is not correctly; expected: %d; actual: %d;", 11, index)
	}
}

func TestStorage_Free(t *testing.T) {
	s := Storage{ lock: sync.RWMutex{} }
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)

	s.Free()
	if s.head != nil {
    t.Errorf("first pointer of nodes is not correctly; expected: %s; actual: %s;", "nil", "not nil")
	}
	if s.size != 0 {
    t.Errorf("length is not correctly; expected: %d; actual: %d;", 0, s.size)
	}

	s.Free()
	if s.Begin() != nil {
    t.Errorf("first pointer of nodes is not correctly; expected: %s; actual: %s;", "nil", "not nil")
	}
	if s.size != 0 {
    t.Errorf("length is not correctly; expected: %d; actual: %d;", 0, s.size)
	}
}
