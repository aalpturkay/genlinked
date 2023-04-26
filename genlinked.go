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

type node[T interface{}] struct {
	data T
	next *node[T]
}

type LinkedList[T interface{}] struct {
	head   *node[T]
	length int
	lock   sync.Mutex
}

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

func (ll *LinkedList[T]) Get(index int) (T, error) {
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

func (ll *LinkedList[T]) GetFirst() (T, error) {
	var data T

	if ll.IsEmpty() {
		return data, errEmptyList
	}

	data = ll.head.data

	return data, nil
}

func (ll *LinkedList[T]) GetLast() (T, error) {
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

func (ll *LinkedList[T]) Size() int {
	return ll.length
}

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
