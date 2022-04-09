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

func NewKeyedQueue(maxSize int) *KeyedQueue {
	if maxSize <= 0 {
		maxSize = 1
	}
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

func (h *KeyedQueue) Push(key string, value []byte) bool {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.size >= h.maxSize {
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

func (h *KeyedQueue) Size() int {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.size
}
