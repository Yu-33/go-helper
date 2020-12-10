package skip

// Iterator used to get the specified range of data.
// The range is start <= x < boundary, and we allowed the start or boundary is nil.
//
// And also implements interface container.Iterator.
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

// Valid represents whether have more elements.
//
// And also implements interface container.Iterator.
func (iter *Iterator) Valid() bool {
	if iter.node == nil || iter.node == iter.end {
		return false
	}

	return true
}

// Next returns a k/v pair and moved the iterator to the next pair.
// Returns nil if no more elements.
//
// And also implements interface container.Iterator.
func (iter *Iterator) Next() KV {
	if !iter.Valid() {
		return nil
	}

	n := iter.node
	iter.node = iter.node.next[0]

	return n
}
