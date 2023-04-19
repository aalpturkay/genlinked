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

func TestGetLast(t *testing.T) {
	type Mana int
	const maxint = ^uint(0) >> 1
	ll := &LinkedList[Mana]{}

	ll.Add(Mana(18))
	ll.Add(Mana(81))
	ll.Add(Mana(19))
	ll.Add(Mana(maxint))

	lastItem, err := ll.GetLast()
	if err != nil {
		t.Errorf("Err: %v\n", err)
	}
	assert.Equal(t, Mana(maxint), lastItem)
}

func TestGetFirst(t *testing.T) {
	type Warrior struct {
		name  string
		power int
	}

	urundir := Warrior{"urundir", 999}
	aragorn := Warrior{"aragorn", 9999}

	ll := &LinkedList[Warrior]{}

	ll.Add(aragorn)
	ll.Add(urundir)

	firstItem, err := ll.GetFirst()

	if err != nil {
		t.Errorf("Err: %v\n", err)
	}
	assert.Equal(t, aragorn, firstItem)
}
