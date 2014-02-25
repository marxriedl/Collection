package collection

//List type
type List interface {
	// Collection, List = List U Collection
	Collection
	// Get returns the element at the specified position, error if out of bounds
	Get(int) (interface {}, error)
	// Set sets the element at the specified position, error if out of bounds
	Set(int, interface {}) error
	// IndexOf returns the index of the first occurrence of the passed element, -1 if not occurring
	IndexOf(interface {}) int
	// LastIndexOf returns the index of the last occurrence of the passed element, -1 if not occurring
	LastIndexOf(interface {}) int
	// AddBegin adds the specified element to the beginning of the List
	AddBegin(interface {}) bool
	// AddAt adds the passed element at the specified position
	AddAt(interface {}, int) (bool, error)
	// AddAllAt adds the elements of the Collection at the specified position
	AddAllAt(Collection, int) (bool, error)
	// RemoveAt removes the element at the specified position
	RemoveAt(int) (bool, error)
	// SubList returns a sub list from lowIndex, including, to highIndex, excluding, error if out of bounds
	SubList(int, int) (List, error)
	// Slice returns a slice form lowIndex, including, to highIndex, excluding, error if out of bounds
	//Slice(int, int) []interface {}
	// Iter returns an Iterator
	Iter() Iterator
}
