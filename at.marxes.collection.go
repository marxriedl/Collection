package main

import "fmt"
import "at/marxes/collection/list"


type IReader interface {
	Read() bool
}

type Reader struct {

}

func (r *Reader) Read() bool{
	return false
}

type ReadWrite struct {
	IReader
}

func main() {
	rw := &ReadWrite{&Reader{}}
	ll := list.CreateLinkedList()
	ll.Add(3)
	val, err := ll.Get(0)
	if err == nil {
		fmt.Printf("%v",val)
	}
	fmt.Printf("%b",rw.Read())          // => "gray"
	//fmt.Println(rw.w)
	fmt.Printf("Hello world!")
}
