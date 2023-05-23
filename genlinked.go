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
	Data T
	Next *node[T]
	Prev *node[T]
}

// LinkedList is a generic and thread-safe implementation of linked-list data structure.
// It can initialized with any type.
type LinkedList[T any] struct {
	Head   *node[T]
	Tail   *node[T]
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
	ll.Head = nil
	ll.Tail = nil
	ll.length = 0

	return ll
}

// Adds T type data into linked-list.
func (ll *LinkedList[T]) Add(data T) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	newNode := &node[T]{
		Data: data,
		Next: nil,
		Prev: nil,
	}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		newNode.Prev = ll.Tail
		ll.Tail.Next = newNode
		ll.Tail = newNode
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
		ll.Tail.Next = &node[T]{
			Data: data,
			Next: nil,
		}
		ll.Tail = ll.Tail.Next

		return nil
	}

	currN := ll.Head

	for i := 0; i < index; i++ {
		currN = currN.Next
	}

	tempN := currN.Next
	currN.Next = &node[T]{
		Data: data,
		Next: nil,
	}

	currN.Next.Next = tempN

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
			ll.Head = nil
			ll.Tail = nil
		} else {
			ll.Head = ll.Head.Next
		}

	} else {
		currN := ll.Head

		for i := 0; i < index-1; i++ {
			if currN == nil || currN.Next == nil {
				return errIndexOutOfRange
			}
			currN = currN.Next
		}

		if currN.Next != nil {
			currN.Next = currN.Next.Next
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

	currN := ll.Head

	for i := 0; i < index; i++ {
		currN = currN.Next
	}

	return currN.Data, nil
}

// Gets first elem.
func (ll *LinkedList[T]) GetFirst() (T, error) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	var data T

	if ll.IsEmpty() {
		return data, errEmptyList
	}

	data = ll.Head.Data

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

	return ll.Tail.Data, nil
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
		currN := ll.Head

		for {
			if currN.Next == nil {
				info += fmt.Sprintf("%v -> nil", currN.Data)
				break
			}
			info += fmt.Sprintf("%v -> ", currN.Data)
			currN = currN.Next
		}
	}

	return info
}
