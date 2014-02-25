package collection

// Iterator type; An iterator over a Collection.
type Iterator interface {
	// Next returns the next element in the Collection
	Next() interface {}
	// HasNext returns true if the Collection has another element, false otherwise
	HasNext() bool
	// Remove removes the element at the current position of the Iterator
	Remove()
}
