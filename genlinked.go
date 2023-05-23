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
	prev *node[T]
}

// LinkedList is a generic and thread-safe implementation of linked-list data structure.
// It can initialized with any type.
type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
	lock   sync.Mutex
}

// Creates an empty linked-list.
func NewLinkedList[T any]() *LinkedList[T] {
	return new(LinkedList[T]).Initialize()
}

// Creates a linked-list with a collection of items.
func NewLinkedListWithItems[T any](items []T) *LinkedList[T] {
	ll := NewLinkedList[T]()
	for _, v := range items {
		ll.Add(v)
	}
	return ll
}

func (ll *LinkedList[T]) Initialize() *LinkedList[T] {
	ll.head = nil
	ll.tail = nil
	ll.length = 0

	return ll
}

// Adds T type data into linked-list.
func (ll *LinkedList[T]) Add(data T) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	newNode := &node[T]{
		data: data,
		next: nil,
		prev: nil,
	}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.prev = ll.tail
		ll.tail.next = newNode
		ll.tail = newNode
	}

	ll.length++
}

func (ll *LinkedList[T]) InsertAfter(index int, data T) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if index >= ll.Size() {
		return errIndexOutOfRange
	}

	if ll.Size()-1 == index {
		ll.tail.next = &node[T]{
			data: data,
			next: nil,
		}
		ll.tail = ll.tail.next

		return nil
	}

	currN := ll.head

	for i := 0; i < index; i++ {
		currN = currN.next
	}

	tempN := currN.next
	currN.next = &node[T]{
		data: data,
		next: nil,
	}

	currN.next.next = tempN

	return nil
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
			ll.tail = nil
		} else {
			ll.head = ll.head.next
		}

	} else {
		currN := ll.head

		for i := 0; i < index-1; i++ {
			if currN == nil || currN.next == nil {
				return errIndexOutOfRange
			}
			currN = currN.next
		}

		if currN.next != nil {
			currN.next = currN.next.next
		}
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

	return ll.tail.data, nil
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
