package container

// Key represents high-level Key type.
type Key Comparator

// Value represents high-level Value type.
type Value interface{}

// KV defines an interface to describe k/v pair.
type KV interface {
	Key() Key
	Value() Value
}

// Container defines an data container interface.
type Container interface {
	// Len returns the number of elements in the container.
	Len() int

	// Insert inserts the key with value in the container.
	// k and v must not be nil, otherwise it will crash.
	// Returns false if key already exists.
	Insert(k Key, v Value) bool

	// Delete remove and returns the value of the specified key.
	// Returns nil if not found.
	Delete(k Key) (v Value)

	// Search get the value of specified key.
	// Returns nil if not found.
	Search(k Key) (v Value)

	// Iter returns an iterator to get the specified range of data.
	// Range: start <= x < boundary, and you should allowed the start or boundary is nil.
	Iter(start Key, boundary Key) Iterator
}

// Iterator defines an iterator interface to returns ranges data.
type Iterator interface {
	// Valid represents whether have more elements.
	Valid() bool

	// Next returns a k/v pair and moved the iterator to the next pair.
	// Returns nil if no more elements.
	Next() KV
}
