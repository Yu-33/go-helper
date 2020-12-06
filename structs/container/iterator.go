package container

type Iterator interface {
	// Valid indicates whether have more elements.
	Valid() bool
	// Next returns an element.
	Next() Comparer
}
