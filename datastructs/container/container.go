package container

type Container interface {
	// Len return the number of elements in container
	Len() int
	// Insert for insert into specified elements, return false if duplicate;
	Insert(elements Comparer) bool
	// Delete for delete specified elements, return nil if not found
	Delete(elements Comparer) Comparer
	// Search for search the specified elements
	Search(elements Comparer) Comparer
	// Iter return a Iterator, include elements: start <= k <= boundary
	Iter(start Comparer, boundary Comparer) Iterator
}
