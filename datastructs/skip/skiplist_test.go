package skip

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/datastructs/container"
)

func output(sl *List) {
	fmt.Println("---------- output ---------")
	for i := 0; i <= sl.level; i++ {
		fmt.Printf("Level <%d, %d> | ", i, sl.lens[i])
		p := sl.head.next[i]
		for p != nil {
			fmt.Printf("%d -> ", p.element)
			p = p.next[i]
		}
		fmt.Printf("\n")
	}
	fmt.Println("---------- output ---------")
}

func checkCorrect(t *testing.T, sl *List) {
	for i := 0; i <= sl.level; i++ {
		p := sl.head.next[i]
		for p != nil && p.next[i] != nil {
			require.Equal(t, p.element.Compare(p.next[i].element), -1)
			p = p.next[i]
		}
	}
}

func TestNew(t *testing.T) {
	sl := New()

	require.NotNil(t, sl)
	require.Equal(t, sl.level, 0)
	require.NotNil(t, sl.r)
	require.NotNil(t, sl.head)
	require.Equal(t, len(sl.lens), maxLevel+1)

	for i := 0; i <= maxLevel; i++ {
		require.Nil(t, sl.head.next[i])
		require.Equal(t, sl.lens[i], 0)
	}
}

func TestSkipList(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	sl := New()

	length := 259
	maxKey := length * 10
	keys := make([]container.Int64, length)

	for x := 0; x < 2; x++ {
		// Test insert
		for i := 0; i < length; i++ {
			for {
				k := container.Int64(r.Intn(maxKey) + 1)
				if ok := sl.Insert(k); ok {
					keys[i] = k
					break
				}
			}
		}

		output(sl)

		require.Equal(t, sl.lens[0], length)
		require.LessOrEqual(t, sl.level, maxLevel)
		checkCorrect(t, sl)

		// boundary
		for _, k := range []container.Int64{0, 0xfffffff} {
			require.True(t, sl.Insert(k))
			require.False(t, sl.Insert(k))
			require.NotNil(t, sl.Search(k))
			require.Equal(t, sl.Search(k).Compare(k), 0)
			require.NotNil(t, sl.Delete(k))
			require.Nil(t, sl.Delete(k))
		}

		// Test get
		for i := 0; i < length; i++ {
			k := keys[i]
			node := sl.Search(k)
			require.NotNil(t, node)
			require.Equal(t, node.Compare(k), 0)
		}

		// Test delete
		for i := 0; i < length; i++ {
			k := keys[i]
			require.NotNil(t, sl.Delete(k))
			require.Nil(t, sl.Delete(k))
		}

		output(sl)
		checkCorrect(t, sl)

		for i := 0; i <= maxLevel; i++ {
			require.Nil(t, sl.head.next[i])
			require.Equal(t, sl.lens[i], 0)
		}
	}
}

func TestList_Search(t *testing.T) {
	sl := New()

	// [24, 61, 67, 84, 91, 130, 133, 145, 150]
	sl.Insert(container.Int64(24))
	sl.Insert(container.Int64(61))
	sl.Insert(container.Int64(67))
	sl.Insert(container.Int64(84))
	sl.Insert(container.Int64(91))
	sl.Insert(container.Int64(130))
	sl.Insert(container.Int64(133))
	sl.Insert(container.Int64(145))
	sl.Insert(container.Int64(150))

	output(sl)

	var node *Node

	// --------- [24, 61, 67, 84, 91, 130, 133, 145, 150] ---------
	node = sl.searchLastLT(container.Int64(21))
	require.Nil(t, node)

	node = sl.searchLastLT(container.Int64(24))
	require.Nil(t, node)

	node = sl.searchLastLT(container.Int64(25))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(24))

	node = sl.searchLastLT(container.Int64(77))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(67))

	node = sl.searchLastLT(container.Int64(132))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(130))

	node = sl.searchLastLT(container.Int64(133))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(130))

	node = sl.searchLastLT(container.Int64(146))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(145))

	node = sl.searchLastLT(container.Int64(150))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(145))

	node = sl.searchLastLT(container.Int64(156))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(150))

	// --------- [24, 61, 67, 84, 91, 130, 133, 145, 150] ---------
	node = sl.searchLastLE(container.Int64(21))
	require.Nil(t, node)

	node = sl.searchLastLE(container.Int64(24))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(24))

	node = sl.searchLastLE(container.Int64(77))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(67))

	node = sl.searchLastLE(container.Int64(132))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(130))

	node = sl.searchLastLE(container.Int64(133))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(133))

	node = sl.searchLastLE(container.Int64(137))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(133))

	node = sl.searchLastLE(container.Int64(150))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(150))

	node = sl.searchLastLE(container.Int64(156))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(150))

	// --------- [24, 61, 67, 84, 91, 130, 133, 145, 150] ---------
	node = sl.searchFirstGT(container.Int64(21))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(24))

	node = sl.searchFirstGT(container.Int64(24))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(61))

	node = sl.searchFirstGT(container.Int64(25))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(61))

	node = sl.searchFirstGT(container.Int64(77))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(84))

	node = sl.searchFirstGT(container.Int64(132))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(133))

	node = sl.searchFirstGT(container.Int64(133))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(145))

	node = sl.searchFirstGT(container.Int64(150))
	require.Nil(t, node)
	node = sl.searchFirstGT(container.Int64(151))
	require.Nil(t, node)

	// --------- [24, 61, 67, 84, 91, 130, 133, 145, 150] ---------
	node = sl.searchFirstGE(container.Int64(21))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(24))

	node = sl.searchFirstGE(container.Int64(24))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(24))

	node = sl.searchFirstGE(container.Int64(25))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(61))

	node = sl.searchFirstGE(container.Int64(77))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(84))

	node = sl.searchFirstGE(container.Int64(132))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(133))

	node = sl.searchFirstGE(container.Int64(133))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(133))

	node = sl.searchFirstGE(container.Int64(146))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(150))

	node = sl.searchFirstGE(container.Int64(150))
	require.NotNil(t, node)
	require.Equal(t, node.element, container.Int64(150))

	node = sl.searchFirstGE(container.Int64(151))
	require.Nil(t, node)

}

func TestList_Iter(t *testing.T) {
	sl := New()

	// --------- [22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150] ---------
	sl.Insert(container.Int64(24))
	sl.Insert(container.Int64(61))
	sl.Insert(container.Int64(67))
	sl.Insert(container.Int64(84))
	sl.Insert(container.Int64(91))
	sl.Insert(container.Int64(130))
	sl.Insert(container.Int64(133))
	sl.Insert(container.Int64(145))
	sl.Insert(container.Int64(150))
	sl.Insert(container.Int64(87))
	sl.Insert(container.Int64(97))
	sl.Insert(container.Int64(22))
	sl.Insert(container.Int64(35))
	sl.Insert(container.Int64(64))
	sl.Insert(container.Int64(76))

	output(sl)

	seeds := []container.Int64{22, 24, 35, 61, 64, 67, 76, 84, 87, 91, 97, 130, 133, 145, 150}

	var iter container.Iterator

	/* ------ test start == nil && boundary == nil */

	iter = sl.Iter(nil, nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	/* ---  test start != nil && boundary == nil --- */
	//
	// start < first node
	iter = sl.Iter(container.Int64(21), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// start == first node
	iter = sl.Iter(container.Int64(22), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// start > first node && start < last node
	iter = sl.Iter(container.Int64(27), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 2; i < len(seeds); i++ {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// start > first node && start < last node
	iter = sl.Iter(container.Int64(62), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 4; i < len(seeds); i++ {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// start > root node && start < last node
	iter = sl.Iter(container.Int64(132), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 12; i < len(seeds); i++ {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// start == last node
	iter = sl.Iter(container.Int64(150), nil)
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	require.Equal(t, iter.Next().(container.Int64), container.Int64(150))
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// start > last node
	iter = sl.Iter(container.Int64(156), nil)
	require.NotNil(t, iter)
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	/* ---  test start == nil && boundary != nil --- */
	//
	// boundary < first node
	iter = sl.Iter(nil, container.Int64(21))
	require.NotNil(t, iter)
	require.False(t, iter.Valid())
	require.Nil(t, iter.Next())

	// boundary == first node
	iter = sl.Iter(nil, container.Int64(22))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	require.Equal(t, iter.Next().(container.Int64), container.Int64(22))
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// boundary < last node && bound > first node
	iter = sl.Iter(nil, container.Int64(147))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds)-1; i++ {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// boundary == last node
	iter = sl.Iter(nil, container.Int64(150))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// boundary > last node
	iter = sl.Iter(nil, container.Int64(156))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := range seeds {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	/* ---  test start != nil && boundary != nil --- */
	//

	// start < boundary && start > first node && bound < last node
	iter = sl.Iter(container.Int64(68), container.Int64(132))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 6; i < len(seeds)-3; i++ {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// start < boundary && start < first node && bound > last node
	iter = sl.Iter(container.Int64(21), container.Int64(153))
	require.NotNil(t, iter)
	require.True(t, iter.Valid())
	for i := 0; i < len(seeds); i++ {
		el := iter.Next()
		require.NotNil(t, el, fmt.Sprintf("%v not found", seeds[i]))
		require.Equal(t, el.(container.Int64), seeds[i])
	}
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// start < boundary && start < first node && bound < first node
	iter = sl.Iter(container.Int64(21), container.Int64(13))
	require.NotNil(t, iter)
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())

	// start > boundary && start > first node
	iter = sl.Iter(container.Int64(65), container.Int64(27))
	require.NotNil(t, iter)
	require.Nil(t, iter.Next())
	require.False(t, iter.Valid())
}
