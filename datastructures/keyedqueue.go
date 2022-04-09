package datastructures

import "sync"

// KeyedQueue is a thread-safe queue that allows for the retrieval of
// elements by key. All elements are unique by key.
type KeyedQueue struct {
	head    *node
	tail    *node
	size    int
	maxSize int
	lock    sync.RWMutex
}

type node struct {
	key   string
	value []byte
	next  *node
}

// NewKeyedQueue creates a new KeyedQueue with the specified maximum size.
// If maxSize is 0, the queue can grow without bound.
func NewKeyedQueue(maxSize int) *KeyedQueue {
	return &KeyedQueue{
		size:    0,
		maxSize: maxSize,
	}
}

// exists returns true if the queue contains an element with the given key.
// If you need to call it from multiple goroutines, depending on the situation
// you will need ensure that you have a lock on the queue.
func (h *KeyedQueue) exists(key string) bool {
	for n := h.head; n != nil; n = n.next {
		if n.key == key {
			return true
		}
	}
	return false
}

// Push adds an element to the queue. If the queue is full, the function
// returns false (the call is non-blocking, and oldest element is not
// removed as it would be in a circular queue implementation).
func (h *KeyedQueue) Push(key string, value []byte) bool {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.size >= h.maxSize && h.maxSize != 0 {
		return false
	}

	if h.exists(key) {
		return true
	}

	node := &node{key: key, value: value}

	if h.head == nil {
		h.head = node
		h.tail = node
	} else {
		h.tail.next = node
		h.tail = node
	}
	h.size++

	return true
}

// Pop removes the oldest element from the queue. If the queue is empty,
// the function returns an empty key string and a nil value.
func (h *KeyedQueue) Pop() (key string, value []byte) {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.head == nil {
		return "", nil
	}

	node := h.head
	h.head = node.next
	h.size--

	return node.key, node.value
}

// Peek returns the value of the element with the given key without removing
// it from the queue. If the queue is empty, the function returns an empty
// key string and a nil value.
func (h *KeyedQueue) Peek(searchKey string) (key string, value []byte) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	for n := h.head; n != nil; n = n.next {
		if n.key == searchKey {
			return n.key, n.value
		}
	}

	return "", nil
}

// Size returns the number of elements in the queue. It is not thread safe but
// thread safety is usually not required for this function.
func (h *KeyedQueue) Size() int {
	return h.size
}
