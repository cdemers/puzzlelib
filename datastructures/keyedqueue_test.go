package datastructures

import "testing"

func TestNewKeyedQueue(t *testing.T) {
	kq := NewKeyedQueue(3)
	if kq == nil {
		t.Errorf("NewKeyedQueue() returned nil")
	}
	if kq.Size() != 0 {
		t.Errorf("KeyedQueue_Size() returned %d, expected 0", kq.Size())
	}
}

func TestKeyedQueue_Push(t *testing.T) {
	kq := NewKeyedQueue(3)
	kq.Push("a", []byte("a"))
	kq.Push("b", []byte("b"))

	if kq.Size() != 2 {
		t.Errorf("KeyedQueue_Push() Size() returned %d, expected 2", kq.Size())
	}
}

func TestKeyedQueue_Pop(t *testing.T) {
	kq := NewKeyedQueue(3)
	kq.Push("a", []byte("1"))
	kq.Push("b", []byte("2"))
	kq.Push("c", []byte("3"))

	if kq.Size() != 3 {
		t.Errorf("KeyedQueue_Size() returned %d, expected 3", kq.Size())
	}

	k, v := kq.Pop()
	if k != "a" {
		t.Errorf("KeyedQueue_Pop() returned %s, expected a", k)
	}
	if string(v) != "1" {
		t.Errorf("KeyedQueue_Pop() returned %s, expected 1", v)
	}

	k, v = kq.Pop()
	if k != "b" {
		t.Errorf("KeyedQueue_Pop() returned %s, expected b", k)
	}
	if string(v) != "2" {
		t.Errorf("KeyedQueue_Pop() returned %s, expected 2", v)
	}

	k, v = kq.Pop()
	if k != "c" {
		t.Errorf("KeyedQueue_Pop() returned %s, expected c", k)
	}
	if string(v) != "3" {
		t.Errorf("KeyedQueue_Pop() returned %s, expected 3", v)
	}

	if kq.Size() != 0 {
		t.Errorf("KeyedQueue_Size() returned %d, expected 0", kq.Size())
	}

	k, v = kq.Pop()
	if k != "" {
		t.Errorf("KeyedQueue_Pop() returned %s, expected empty string", k)
	}
	if v != nil {
		t.Errorf("KeyedQueue_Pop() returned %s, expected nil", v)
	}
}

func TestKeyedQueue_Push2(t *testing.T) {
	kq := NewKeyedQueue(3)
	kq.Push("a", []byte("1"))
	kq.Push("b", []byte("2"))
	kq.Push("c", []byte("3"))

	if kq.Size() != 3 {
		t.Errorf("KeyedQueue_Size() returned %d, expected 3", kq.Size())
	}

	kq.Push("b", []byte("2"))

	if kq.Size() != 3 {
		t.Errorf("KeyedQueue_Size() returned %d, expected 3", kq.Size())
	}
}

func TestKeyedQueue_Peek(t *testing.T) {
	kb := NewKeyedQueue(3)
	kb.Push("a", []byte("1"))
	kb.Push("b", []byte("2"))
	kb.Push("c", []byte("3"))

	k, v := kb.Peek("b")
	if k != "b" {
		t.Errorf("KeyedQueue_Peek() returned %s, expected b", k)
	}
	if string(v) != "2" {
		t.Errorf("KeyedQueue_Peek() returned %s, expected 2", v)
	}

	if kb.Size() != 3 {
		t.Errorf("KeyedQueue_Size() returned %d, expected 3", kb.Size())
	}
}
