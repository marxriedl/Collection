package list

import "testing"

func TestEmpty(t *testing.T) {
	ll := CreateLinkedList()
	e := ll.Empty()
	if !e {
		t.Fail()
	}
}

func TestNonEmpty(t *testing.T) {
	ll := CreateLinkedList()
	ll.Add(3)
	e := ll.Empty()
	if e {
		t.Fail()
	}
}


func TestSize(t *testing.T) {
	ll := CreateLinkedList()
	s := ll.Size()
	if s != 0 {
		t.Fail()
	}
}

func TestSizeElement(t *testing.T) {
	ll := CreateLinkedList()
	ll.Add(3)
	s := ll.Size()
	if s != 1 {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	ll := CreateLinkedList()
	val := ll.Add(3)
	if !val {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	ll := CreateLinkedList()
	val := ll.Add(3)
	contains := ll.Contains(3)
	if !val && !contains {
		t.Fail()
	}
}

func TestContainsMultipleElements(t *testing.T) {
	ll := CreateLinkedList()
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	contains := ll.Contains(1)
	if contains {
		t.Fail()
	}
}

func TestNotContains(t *testing.T) {
	ll := CreateLinkedList()
	ll.Add(3)
	ll.Add(4)
	contains := ll.Contains(5)
	if contains {
		t.Fail()
	}
}

func TestNotContainsEmptyColl(t *testing.T) {
	ll := CreateLinkedList()
	contains := ll.Contains(3)
	if contains {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	ll := CreateLinkedList()
	val, err := ll.Get(0)
	if val != nil || err ==nil {
		t.Fail()
	}
}

func TestGetRealValue(t *testing.T) {
	ll := CreateLinkedList()
	ll.Add(3)
	val, err := ll.Get(0)
	if val != 3 || err != nil {
		t.Fail()
	}
}

func TestContainsGetMultipleElements(t *testing.T) {
	ll := CreateLinkedList()
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	val0, err0 := ll.Get(0)
	val1, err1 := ll.Get(1)
	val2, err2 := ll.Get(2)
	if val0 != 3 || err0 !=nil {
		t.Errorf("expected %i was %i", 3, val0)
		t.Fail()
	}
	if val1 != 4 || err1 !=nil {
		t.Errorf("expected %i was %i", 4, val1)
		t.Fail()
	}
	if val2 != 5 || err2 !=nil{
		t.Errorf("expected %i was %i", 5, val2)
		t.Fail()
	}
}

func TestIterChan(t *testing.T) {
	ll := CreateLinkedList()
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)
	go func() {
		val := 3
		for data := range ll.IterChan(){
			if data == val {
				t.Fail()
			}
			val++
		}
	}();
}
