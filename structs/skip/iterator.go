package skip

// Iterator to get the data for the specified range.
//
// Range interval: ( start <= x < boundary ).
// We will return data from the beginning if start is nil,
// And return data util the end if boundary is nil.
//
// And it is also the implementation of interface container.Iterator
type Iterator struct {
	node *listNode
	end  *listNode
}

// creates an Iterator.
func newIterator(sl *List, start Key, boundary Key) *Iterator {
	var node, end *listNode

	// If both the start and boundary are not nil, the start should less than the boundary.
	if !(start != nil && boundary != nil && start.Compare(boundary) != -1) {
		if start == nil {
			node = sl.head.next[0]
		} else {
			node = sl.searchFirstGE(start)
		}

		if boundary != nil {
			end = sl.searchFirstGE(boundary)
		}
	}

	iter := &Iterator{
		node: node,
		end:  end,
	}
	return iter
}

// Valid represents whether to have more elements in the Iterator.
func (iter *Iterator) Valid() bool {
	if iter.node == nil || iter.node == iter.end {
		return false
	}
	return true
}

// Next returns a k/v pair and moved the iterator to the next pair.
// Returns nil if no more elements.
func (iter *Iterator) Next() KV {
	if !iter.Valid() {
		return nil
	}

	n := iter.node
	iter.node = iter.node.next[0]

	return n
}
