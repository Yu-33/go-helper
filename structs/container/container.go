package container

// Key represents high-level Key type.
type Key = Comparator

// Value represents high-level Value type.
type Value interface{}

// KV defines an interface to describe k/v pair.
type KV interface {
	// Key returns the key.
	Key() Key
	// Value returns the value.
	Value() Value
}

// Container defines an data container interface.
type Container interface {
	// Len returns the number of elements.
	Len() int

	// Insert inserts the giving key and value.
	// Returns false if key already exists.
	Insert(k Key, v Value) bool

	// Delete removes and returns the KV structure corresponding to the given key.
	// Returns nil if not found.
	Delete(k Key) (kv KV)

	// Search returns the KV structure corresponding to the given key.
	// Returns nil if not found.
	Search(k Key) (kv KV)

	// Iter creates an iterator to get the data for the specified range.
	//
	// Range interval: ( start <= x < boundary ).
	// We will return data from the beginning if start is nil,
	// And return data util the end if boundary is nil.
	Iter(start Key, boundary Key) Iterator
}

// Iterator defines an iterator interface to returns multiple data.
type Iterator interface {
	// Valid represents whether to have more elements in the Iterator.
	Valid() bool

	// Next returns a k/v pair and moved the iterator to the next pair.
	// Returns nil if no more elements.
	Next() KV
}
