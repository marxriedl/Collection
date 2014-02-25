package collection

//Collection interface, root interface of the collection hierarchy
type Collection interface {
	// Clear empties the collection
	Clear()
	// Empty returns true if the collection is empty, otherwise false
	Empty() bool
	// Size returns the current size of the Collection
	Size() int
	// Contains returns true if the collection is empty, otherwise false
	Contains(interface {}) bool
	// ContainsAll returns true if the elements in the passed Collection all are found, otherwise false
	ContainsAll(*Collection) bool
	// Add returns true if the element can be added to the Collection, otherwise false
	Add(interface {}) bool
	// AddAll adds all elements of the passed Collection to this Collection
	AddAll(*Collection) bool
	// Remove returns true if the element can be removed to the Collection, otherwise false
	Remove(interface {}) bool
	// RemoveAll removes all elements in the Collection that are common to both Collections
	RemoveAll(*Collection) bool
	// Intersect retains all elements in the Collection that are common to both Collections
	Intersect(*Collection) bool
	// Slice returns a slice of the collection
	Slice() []interface {}
	//Apply applies the provided function to all elements in the Collection, returns true if successful, otherwise false
	Apply(func(interface {}) interface {}) bool
	//Collect collects all elements of the Collection that fulfill the specified function
	Collect(func(interface {}) bool) Collection
	//IterChan returns a channel to iterate the Collection
	IterChan() <- chan interface {}
}
