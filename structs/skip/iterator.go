package skip

type Iterator struct {
	node *Node
	end  *Node
}

func newIterator(sl *List, start Element, boundary Element) *Iterator {
	var node, end *Node

	if !(start != nil && boundary != nil && start.Compare(boundary) == 1) {
		if start == nil {
			node = sl.head.next[0]
		} else {
			node = sl.searchFirstGE(start)
		}

		if boundary != nil {
			end = sl.searchFirstGT(boundary)
		}
	}

	iter := &Iterator{
		node: node,
		end:  end,
	}
	return iter
}

func (iter *Iterator) Valid() bool {
	if iter.node == nil || iter.node == iter.end {
		return false
	}

	return true
}

func (iter *Iterator) Next() Element {
	if !iter.Valid() {
		return nil
	}

	elements := iter.node.element
	iter.node = iter.node.next[0]

	return elements
}
