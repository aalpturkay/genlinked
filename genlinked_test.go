package genlinked

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[bool]()

	ll.Add(true)

	assert.Equal(t, 1, ll.Size())
}

func TestNewLinkedListWithItems(t *testing.T) {
	categories := NewLinkedListWithItems([]string{"burger", "pizza", "wrap", "icecream"})

	assert.Equal(t, 4, categories.Size())
}

func TestAddElemToLinkedList(t *testing.T) {

	ll := NewLinkedList[string]()

	ll.Add("warrior")

	ll.Add("mage")

	ll.Add("shaman")

	ll.Add("barbarian")

	ll.Add("paladin")

	assert.Equal(t, 5, ll.length)
}

func TestGetElemByIndexFromLinkedList(t *testing.T) {

	ll := NewLinkedList[float64]()

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

	ll := NewLinkedList[string]()

	ll.Add("frodo")
	ll.Add("sam")
	ll.Add("gandalf")

	_, err := ll.Get(3)

	assert.Error(t, err, errIndexOutOfRange)
}

func TestGetElemByIndexFromLinkedListWhenListEmpty(t *testing.T) {

	ll := NewLinkedList[string]()

	_, err := ll.Get(0)

	assert.Error(t, err, errIndexOutOfRange)
}

func TestGetSize(t *testing.T) {

	ll := NewLinkedList[string]()

	ll.Add("warrior")

	ll.Add("mage")

	ll.Add("shaman")

	assert.Equal(t, 3, ll.Size())
}

func TestPrintLinkedList(t *testing.T) {
	ll := NewLinkedList[string]()
	ll.Add("warrior")

	ll.Add("mage")

	ll.Add("shaman")

	ll.Add("barbarian")

	ll.Add("paladin")

	fmt.Println(ll)
}

func TestIsLinkedListEmpty(t *testing.T) {
	ll := NewLinkedList[string]()

	assert.Equal(t, true, ll.IsEmpty())
}

func TestGetLast(t *testing.T) {
	type Mana int
	const maxint = ^uint(0) >> 1
	ll := NewLinkedList[Mana]()

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

	ll := NewLinkedList[Warrior]()

	ll.Add(aragorn)
	ll.Add(urundir)

	firstItem, err := ll.GetFirst()

	if err != nil {
		t.Errorf("Err: %v\n", err)
	}
	assert.Equal(t, aragorn, firstItem)
}

func TestString(t *testing.T) {
	movies := NewLinkedList[string]()

	movies.Add("Harry Potter")
	movies.Add("LOTR")

	fmt.Println(movies)
}

func TestIsEmpty(t *testing.T) {
	langs := NewLinkedList[string]()

	langs.Add("ruby")

	empty := langs.IsEmpty()

	assert.False(t, empty)
}

func TestRemove(t *testing.T) {
	series := NewLinkedList[string]()

	series.Add("KV")
	series.Add("Breaking Bad")
	series.Add("GOT")

	series.Remove(1)
	fmt.Println(series)
	assert.Equal(t, 2, series.Size())
}

func TestRemoveWhenIndexOOR(t *testing.T) {
	series := NewLinkedList[string]()

	series.Add("KV")
	series.Add("Breaking Bad")
	series.Add("GOT")

	err := series.Remove(8)
	fmt.Println(series)
	assert.Error(t, errIndexOutOfRange, err)
}

func TestCalAddTime(t *testing.T) {
	startT := time.Now()
	ll := NewLinkedList[int]()

	size := 5_000

	for i := 0; i < size; i++ {
		ll.Add(1)
	}

	endT := time.Now()

	deltaT := endT.Sub(startT)

	fmt.Printf("linked-list deltaT: %v\n", deltaT)
}

func TestInsert(t *testing.T) {
	categories := NewLinkedListWithItems([]string{"burger", "pizza", "wrap", "icecream"})

	err := categories.InsertAfter(3, "döner")
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
		return
	}

	fmt.Println(categories)
}

// it takes: ~42 nanoseconds
func TestInsertAfterTime(t *testing.T) {
	ll := NewLinkedList[int]()

	size := 5_000

	for i := 0; i < size; i++ {
		ll.Add(1)
	}

	startT := time.Now()

	err := ll.InsertAfter(3, 10)
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
		return
	}

	endT := time.Now()

	deltaT := endT.Sub(startT)

	fmt.Printf("linked-list deltaT: %v nsec\n", deltaT.Nanoseconds())
}

// it takes: ~13750 nanoseconds
func TestInsertElemSliceTime(t *testing.T) {
	numbers := make([]int, 0)

	for i := 0; i < 5_000; i++ {
		numbers = append(numbers, 1)
	}
	startT := time.Now()

	newData := 6
	targetIndex := 3

	numbers = append(numbers[:targetIndex+1], append([]int{newData}, numbers[targetIndex+1:]...)...)
	endT := time.Now()

	deltaT := endT.Sub(startT)
	fmt.Printf("slice deltaT: %v nsec\n", deltaT.Nanoseconds())
}

func TestInsertEnfOfTheList(t *testing.T) {
	ll := NewLinkedListWithItems([]string{"aykut", "alp"})
	err := ll.InsertAfter(ll.Size()-1, "turkay")
	if err != nil {
		t.Errorf("err: %s\n", err.Error())
		return
	}

	fmt.Println(ll)

	last, _ := ll.GetLast()

	assert.Equal(t, "turkay", last)
}

func TestThreadSafety(t *testing.T) {
	operations := 1_000

	ll := NewLinkedList[int]()
	fmt.Println(ll.Size())
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("adding")
		defer wg.Done()
		for i := 0; i < operations; i++ {
			ll.Add(i)
		}
	}()

	go func() {
		fmt.Println("removing")
		defer wg.Done()
		for i := 0; i < operations; i++ {
			ll.Remove(i)
		}
	}()

	wg.Wait()

	curr := ll.Head
	for curr != nil {
		if curr.Next != nil && curr.Data >= curr.Next.Data {
			t.Errorf("list must be in sorted order.")
		}
		curr = curr.Next
	}
}
