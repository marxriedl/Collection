package list

import "errors"
import "at/marxes/collection"

type node struct {
	data interface {}
	next *node
}

type LinkedList struct {
	size int
	head *node
	end *node
}

type iterator struct {
	current *node
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{0, nil, nil}
}

func insert(n, in *node) {
	in.next = n.next
	n.next = in
}

func remove(p, re *node) {
	p.next = re.next
	re.next = nil
}

func findValue(ll *LinkedList, val interface {}) (prev, n *node){
	it := ll.head
	p := ll.head
	for it != nil {
		if it.data == val {
			return p, it
		}
		p = it
		it = it.next
	}
	return nil,nil
}

func findIndex(ll *LinkedList, pos int) (prev, n *node){
	it := ll.head
	p := ll.head
	for idx := 0; idx < pos ; idx ++ {
		idx++
		p = it
		it = it.next
	}
	return p, it
}


func (ll *LinkedList) Clear(){
 	ll.head = nil
	ll.end = nil
	ll.size = 0;
}

func (ll *LinkedList) Empty() bool{
	return ll.size == 0
}
func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) Contains(val interface {}) bool {
	p, it := findValue(ll, val)
	if it != nil {return true}
	return false
}

func (ll *LinkedList) ContainsAll(coll *collection.Collection) bool{
	for data := range (*coll).IterChan(){
		if !ll.Contains(data) {return false}
	}
	return true
}

func (ll *LinkedList) AddBegin(val interface {}) bool {
	if ll.head == nil {
		ll.head = &node{val, nil}
		ll.end = ll.head
	} else {
		ll.head = &node{val, ll.head}
	}
	ll.size++
	return true
}
func (ll *LinkedList) Add(val interface {}) bool {
	if ll.head == nil {
		ll.head = &node{val, nil}
		ll.end = ll.head
	} else {
		n := &node{val, nil}
		insert(ll.end, n)
		ll.end = n
	}
	ll.size++
	return true
}

func (ll *LinkedList) AddAll(coll *collection.Collection) bool {
	for data := range (*coll).IterChan() {
		ll.Add(data)
	}
	return true
}

func (ll *LinkedList) Remove(val interface {}) bool {
	p, n := findValue(ll, val)
	if n != nil {
		if p == nil {
			ll.head = n.next
			n.next = nil
		}else{
			remove(p ,n)
		}
		ll.size--
	  	return true
	}
	return false
}

func (ll *LinkedList) RemoveAll(coll *collection.Collection) bool {
	for data := range (*coll).IterChan() {
		ll.Remove(data)
	}
	return true
}
func (ll *LinkedList) Intersect(coll *collection.Collection) bool {
	for data := range ll.IterChan() {
		if !(*coll).Contains(data) {
			ll.Remove(data)
		}
	}
	return false
}
func (ll *LinkedList) Slice() []interface {} {return nil}

func (ll *LinkedList) Apply(f func(interface {}) interface {}) bool {
	it := ll.head
	for it!=nil {
		it.data = f(it.data)
		it = it.next
	}
	return true
}

func (ll *LinkedList) Collect(f func(interface {}) bool) collection.Collection{
	list := CreateLinkedList()
	it := ll.head
	for it!=nil {
		if f(it.data) {
			list.Add(it.data)
		}
		it = it.next
	}
	return list
}

func (ll *LinkedList)  IterChan() <- chan interface {} {
	ch := make(chan interface {});
	go func() {
		it := ll.head
		for it != nil {
			ch <- it.data
			it = it.next
		}
	}();
	return ch
}

func (ll *LinkedList) Get(idx int) (interface {}, error){
	if idx < 0 || idx >= ll.size {
		return nil, errors.New("Index Out of Bounds")
	}
	p, n := findIndex(ll, idx)
	return n.data, nil
}
func (ll *LinkedList) Set(idx int, val interface {}) error {
	if idx < 0 || idx >= ll.size {
		return errors.New("Index Out of Bounds")
	}

	p, n := findIndex(ll, idx)
	insert(n, &node{val, nil})
	return nil
}
func (ll *LinkedList) IndexOf(val interface {}) int {
	it := ll.head
	for i := 0 ; it != nil; i++ {
		if it.data == val {
			return i
		}
		it = it.next
	}
	return -1
}
func (ll *LinkedList) LastIndexOf(val interface {}) int {
	r := -1
	it := ll.head
	for i := 0 ; it != nil; i++ {
		if it.data == val {
			r = i
		}
		it = it.next
	}
	return r
}
func (ll *LinkedList) AddAt(val interface {}, idx int) (bool, error) {
	if idx < 0 || idx >= ll.size {
		return false, errors.New("Index Out of Bounds")
	}
	p, n := findIndex(ll, idx)
	insert(n, &node{val, nil})
	return true, nil
}
func (ll *LinkedList) AddAllAt(coll *collection.Collection, idx int) (bool, error) {
	if idx < 0 || idx >= ll.size {
		return false, errors.New("Index Out of Bounds")
	}
	i := idx
	for data := range (*coll).IterChan() {
		ll.AddAt(data, i)
		i++
	}
	return true, nil
}
func (ll *LinkedList) RemoveAt(idx int) (bool, error) {
	if idx < 0 || idx >= ll.size {
		return false, errors.New("Index Out of Bounds")
	}
	p, n := findIndex(ll, idx)
	remove(p, n)
	return true, nil
}
func (ll *LinkedList) SubList(int, int) (*collection.List, error) {return nil, nil}
func (ll *LinkedList) Iter() collection.Iterator {
	return nil // &iterator{ll.head}
}
