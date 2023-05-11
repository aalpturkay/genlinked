package genlinked

import (
	"errors"
	"fmt"
	"sync"
)

var (
	errIndexOutOfRange error = errors.New("index out of range")
	errEmptyList       error = errors.New("list is empty")
)

type node[T any] struct {
	data T
	next *node[T]
}

// LinkedList is a generic and thread-safe implementation of linked-list data structure.
// It can initialized with any type.
type LinkedList[T any] struct {
	head   *node[T]
	length int
	lock   sync.Mutex
}

// Creates an empty linked-list.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Creates a linked-list with a collection of items.
func NewLinkedListWithItems[T any](items []T) *LinkedList[T] {
	ll := NewLinkedList[T]()
	for _, v := range items {
		ll.Add(v)
	}
	return ll
}

// Adds T type data into linked-list.
func (ll *LinkedList[T]) Add(data T) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	newNode := &node[T]{
		data: data,
		next: nil,
	}

	if ll.head == nil {
		ll.head = newNode

	} else {
		currN := ll.head
		for {
			if currN.next == nil {
				break
			}
			currN = currN.next
		}
		currN.next = newNode
	}

	ll.length++
}

// Removes an elem. from linked-list by given index.
func (ll *LinkedList[T]) Remove(index int) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if index < 0 || index >= ll.length {
		return errIndexOutOfRange
	}

	if index == 0 {
		if ll.length == 1 {
			ll.head = nil
		} else {
			ll.head = ll.head.next
		}

	} else {
		currN := ll.head

		for i := 0; i < index-1; i++ {
			currN = currN.next
		}

		currN.next = currN.next.next
	}

	ll.length--
	return nil
}

// Gets an elem. from linked-list by given index.
func (ll *LinkedList[T]) Get(index int) (T, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	var data T

	if index < 0 || index >= ll.length {
		return data, errIndexOutOfRange
	}

	currN := ll.head

	for i := 0; i < index; i++ {
		currN = currN.next
	}

	return currN.data, nil
}

// Gets first elem.
func (ll *LinkedList[T]) GetFirst() (T, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	var data T

	if ll.IsEmpty() {
		return data, errEmptyList
	}

	data = ll.head.data

	return data, nil
}

// Gets last elem.
func (ll *LinkedList[T]) GetLast() (T, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	var data T

	if ll.IsEmpty() {
		return data, errEmptyList
	}

	currN := ll.head

	for {
		if currN.next == nil {
			break
		}
		currN = currN.next
	}

	return currN.data, nil
}

// Returns length of linked-list.
func (ll *LinkedList[T]) Size() int {
	return ll.length
}

// Returns a bool that indicates whether linked-list size is equal to zero.
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.Size() == 0
}

func (ll *LinkedList[T]) String() string {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	info := ""

	if !ll.IsEmpty() {
		currN := ll.head

		for {
			if currN.next == nil {
				info += fmt.Sprintf("%v -> nil", currN.data)
				break
			}
			info += fmt.Sprintf("%v -> ", currN.data)
			currN = currN.next
		}
	}

	return info
}
