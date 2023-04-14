package genlinked

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddElemToLinkedList(t *testing.T) {

	ll := &LinkedList[string]{}

	ll.Add("warrior")

	ll.Add("mage")

	ll.Add("shaman")

	ll.Add("barbarian")

	ll.Add("paladin")

	assert.Equal(t, 5, ll.length)
}

func TestGetElemByIndexFromLinkedList(t *testing.T) {

	ll := &LinkedList[float64]{}

	ll.Add(3)
	ll.Add(1.1)
	ll.Add(2.3)

	val, err := ll.Get(2)
	if err != nil {
		t.Errorf("Err: %v\n", err)
	}

	assert.Equal(t, val, 2.3)
}

func TestGetElemByIndexWhenIndexOutOfRangeFromLinkedList(t *testing.T) {

	ll := &LinkedList[string]{}

	ll.Add("frodo")
	ll.Add("sam")
	ll.Add("gandalf")

	_, err := ll.Get(3)

	assert.Error(t, err, errIndexOutOfRange)
}

func TestGetElemByIndexFromLinkedListWhenListEmpty(t *testing.T) {

	ll := &LinkedList[string]{}

	_, err := ll.Get(0)

	assert.Error(t, err, errIndexOutOfRange)
}

func TestGetSize(t *testing.T) {

	ll := &LinkedList[string]{}

	ll.Add("warrior")

	ll.Add("mage")

	ll.Add("shaman")

	assert.Equal(t, 3, ll.Size())
}

func TestPrintLinkedList(t *testing.T) {
	ll := &LinkedList[string]{}

	ll.Add("warrior")

	ll.Add("mage")

	ll.Add("shaman")

	ll.Add("barbarian")

	ll.Add("paladin")

	fmt.Println(ll)
}

func TestIsLinkedListEmpty(t *testing.T) {
	ll := &LinkedList[string]{}

	assert.Equal(t, true, ll.IsEmpty())
}
