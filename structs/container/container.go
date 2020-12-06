package container

type Container interface {
	// Len return the number of elements in container
	Len() int
	// Insert for insert into specified element, return false if duplicate;
	Insert(element Comparer) bool
	// Delete for delete specified element, return nil if not found
	Delete(element Comparer) Comparer
	// Search for search the specified element
	Search(element Comparer) Comparer
	// Iter return a Iterator, include element: start <= k <= boundary
	Iter(start Comparer, boundary Comparer) Iterator
}
