package list

import "errors"
import "at/marxes/collection"

type node struct {
	data interface {}
	next *node
}

type linkedlist struct {
	size int
	head *node
	tail *node
}

type iterator struct {
	current *node
	ll *linkedlist
}


func (it *iterator) Next() interface {} {
	if it.HasNext() {
		it.current = it.current.next
		return it.current.data
	}
	return nil
}

func (it *iterator) HasNext() bool {
	return it.current.next != nil
}


func (it *iterator)  Remove() {
	val := it.current.data
	it.current = it.current.next
	it.ll.Remove(val)
}

func CreateLinkedList() collection.List {
	return &linkedlist{0, nil, nil}
}

func insert(n, in *node) {
	in.next = n.next
	n.next = in
}

func remove(p, re *node) {
	p.next = re.next
	re.next = nil
}

func findValue(ll *linkedlist, val interface {}) (prev, n *node){
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

func findIndexP(ll *linkedlist, pos int) (prev, n *node){
	it := ll.head
	p := ll.head
	for idx := 0; idx < pos ; idx ++ {
		idx++
		p = it
		it = it.next
	}
	return p, it
}

func findIndex(ll *linkedlist, pos int) (n *node){
	it := ll.head
	for idx := 0; idx < pos ; idx ++ {
		idx++
		it = it.next
	}
	return it
}


func (ll *linkedlist) Clear(){
 	ll.head = nil
	ll.tail = nil
	ll.size = 0;
}

func (ll *linkedlist) Empty() bool{
	return ll.size == 0
}
func (ll *linkedlist) Size() int {
	return ll.size
}

func (ll *linkedlist) Contains(val interface {}) bool {
	p, it := findValue(ll, val)
	if (p != nil || p == nil) && it != nil {return true}
	return false
}

func (ll *linkedlist) ContainsAll(coll collection.Collection) bool{
	for data := range coll.IterChan(){
		if !ll.Contains(data) {return false}
	}
	return true
}

func (ll *linkedlist) AddBegin(val interface {}) bool {
	if ll.head == nil {
		ll.head = &node{val, nil}
		ll.tail = ll.head
	} else {
		ll.head = &node{val, ll.head}
	}
	ll.size++
	return true
}
func (ll *linkedlist) Add(val interface {}) bool {
	if ll.head == nil {
		ll.head = &node{val, nil}
		ll.tail = ll.head
	} else {
		n := &node{val, nil}
		insert(ll.tail, n)
		ll.tail = n
	}
	ll.size++
	return true
}

func (ll *linkedlist) AddAll(coll collection.Collection) bool {
	for data := range coll.IterChan() {
		ll.Add(data)
	}
	return true
}

func (ll *linkedlist) Remove(val interface {}) bool {
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

func (ll *linkedlist) RemoveAll(coll collection.Collection) bool {
	for data := range coll.IterChan() {
		ll.Remove(data)
	}
	return true
}
func (ll *linkedlist) Intersect(coll collection.Collection) bool {
	for data := range ll.IterChan() {
		if !coll.Contains(data) {
			ll.Remove(data)
		}
	}
	return false
}
func (ll *linkedlist) Slice() []interface {} {
	arr := make([]interface {}, ll.Size())
	i := 0
	for val := range ll.IterChan() {
		arr[i] = val
		i++
	}
	return arr
}

func (ll *linkedlist) Apply(f func(interface {}) interface {}) bool {
	it := ll.head
	for it!=nil {
		it.data = f(it.data)
		it = it.next
	}
	return true
}

func (ll *linkedlist) Collect(f func(interface {}) bool) collection.Collection{
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

func (ll *linkedlist)  IterChan() <- chan interface {} {
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

func (ll *linkedlist) Get(idx int) (interface {}, error){
	if idx < 0 || idx >= ll.size {
		return nil, errors.New("Index Out of Bounds")
	}
	n := findIndex(ll, idx)
	return n.data, nil
}
func (ll *linkedlist) Set(idx int, val interface {}) error {
	if idx < 0 || idx >= ll.size {
		return errors.New("Index Out of Bounds")
	}

	n := findIndex(ll, idx)
	insert(n, &node{val, nil})
	return nil
}
func (ll *linkedlist) IndexOf(val interface {}) int {
	it := ll.head
	for i := 0 ; it != nil; i++ {
		if it.data == val {
			return i
		}
		it = it.next
	}
	return -1
}
func (ll *linkedlist) LastIndexOf(val interface {}) int {
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
func (ll *linkedlist) AddAt(val interface {}, idx int) (bool, error) {
	if idx < 0 || idx >= ll.size {
		return false, errors.New("Index Out of Bounds")
	}
	n := findIndex(ll, idx)
	insert(n, &node{val, nil})
	return true, nil
}
func (ll *linkedlist) AddAllAt(coll collection.Collection, idx int) (bool, error) {
	if idx < 0 || idx >= ll.size {
		return false, errors.New("Index Out of Bounds")
	}
	i := idx
	for data := range coll.IterChan() {
		ll.AddAt(data, i)
		i++
	}
	return true, nil
}
func (ll *linkedlist) RemoveAt(idx int) (bool, error) {
	if idx < 0 || idx >= ll.size {
		return false, errors.New("Index Out of Bounds")
	}
	p, n := findIndexP(ll, idx)
	remove(p, n)
	return true, nil
}
func (ll *linkedlist) SubList(low, high int) (collection.List, error) {
	if low < 0 || low >= ll.size || high < 0 || high >= ll.size || low > high {
		return nil, errors.New("Indexes Out of Bounds")
	}
	list := CreateLinkedList()
	for i := low; i < high; i++{
		val, err := ll.Get(i)
		if err == nil {
			list.Add(val)
		}

	}
	return list, nil
}
func (ll *linkedlist) Iter() collection.Iterator {
	return &iterator{ll.head, ll}
}
