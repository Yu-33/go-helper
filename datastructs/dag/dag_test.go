package dag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/helper/datastructs/container"
)

func TestNew(t *testing.T) {
	g := New()
	require.NotNil(t, g)
	require.NotNil(t, g.vertexes)
}

func TestDAG_AddVertex(t *testing.T) {
	g := New()
	vertexes := []container.Int64{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// Add vertex
	for i := 0; i < len(vertexes); i++ {
		require.True(t, g.AddVertex(vertexes[i]))
		require.False(t, g.AddVertex(vertexes[i]))
	}

	// test internal elements
	it1 := g.vertexes.Iter(nil, nil)
	for it1.Valid() {
		elements := it1.Next()
		require.NotNil(t, elements)
		n, ok := elements.(*node)
		require.True(t, ok)
		require.NotNil(t, n.vex)
		require.NotNil(t, n.in)
		require.NotNil(t, n.out)
	}
}

func TestDAG_AddEdge(t *testing.T) {
	g := New()
	vertexes := []container.Int64{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// Test Add vertex
	for i := 0; i < len(vertexes); i++ {
		require.True(t, g.AddVertex(vertexes[i]))
	}

	// Test Add edge positive
	require.True(t, g.AddEdge(vertexes[0], vertexes[1]))

	require.True(t, g.AddEdge(vertexes[1], vertexes[2]))
	require.True(t, g.AddEdge(vertexes[1], vertexes[3]))
	require.True(t, g.AddEdge(vertexes[1], vertexes[6]))

	require.True(t, g.AddEdge(vertexes[2], vertexes[4]))
	require.True(t, g.AddEdge(vertexes[2], vertexes[5]))
	require.True(t, g.AddEdge(vertexes[2], vertexes[6]))

	require.True(t, g.AddEdge(vertexes[3], vertexes[2]))
	require.True(t, g.AddEdge(vertexes[3], vertexes[6]))
	require.True(t, g.AddEdge(vertexes[3], vertexes[7]))

	require.True(t, g.AddEdge(vertexes[4], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[5], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[6], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[7], vertexes[8]))

	fmt.Println("g.OutDegreeToString():", g.OutDegreeToString())
	fmt.Println("g.InDegreeToString():", g.InDegreeToString())

	// Test Add edge negative
	require.False(t, g.AddEdge(vertexes[1], vertexes[1]))
	require.False(t, g.AddEdge(vertexes[6], vertexes[1]))
	require.False(t, g.AddEdge(vertexes[2], vertexes[1]))
	require.False(t, g.AddEdge(vertexes[8], vertexes[3]))
}

func TestDAG_DelVertex(t *testing.T) {
	g := New()
	vertexes := []container.Int64{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// Test Add vertex
	for i := 0; i < len(vertexes); i++ {
		require.True(t, g.AddVertex(vertexes[i]))
	}

	// Test Add edge positive
	require.True(t, g.AddEdge(vertexes[0], vertexes[1]))

	require.True(t, g.AddEdge(vertexes[1], vertexes[2]))
	require.True(t, g.AddEdge(vertexes[1], vertexes[3]))
	require.True(t, g.AddEdge(vertexes[1], vertexes[6]))

	require.True(t, g.AddEdge(vertexes[2], vertexes[4]))
	require.True(t, g.AddEdge(vertexes[2], vertexes[5]))
	require.True(t, g.AddEdge(vertexes[2], vertexes[6]))

	require.True(t, g.AddEdge(vertexes[3], vertexes[2]))
	require.True(t, g.AddEdge(vertexes[3], vertexes[6]))
	require.True(t, g.AddEdge(vertexes[3], vertexes[7]))

	require.True(t, g.AddEdge(vertexes[4], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[5], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[6], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[7], vertexes[8]))

	// Test Delete vertex
	require.True(t, g.DelVertex(vertexes[3]))
	require.False(t, g.DelVertex(vertexes[3]))

	it1 := g.vertexes.Iter(nil, nil)
	for it1.Valid() {
		n1 := it1.Next().(*node)
		require.NotEqual(t, n1.vex, vertexes[3])

		it2 := n1.in.Iter(nil, nil)
		for it2.Valid() {
			require.NotEqual(t, it2.Next().(*node).vex, vertexes[3])
		}
		it3 := n1.out.Iter(nil, nil)
		for it3.Valid() {
			require.NotEqual(t, it3.Next().(*node).vex, vertexes[3])
		}
	}
}


func TestDAG_DelEdge(t *testing.T) {
	g := New()
	vertexes := []container.Int64{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// Test Add vertex
	for i := 0; i < len(vertexes); i++ {
		require.True(t, g.AddVertex(vertexes[i]))
	}

	// Test Add edge positive
	require.True(t, g.AddEdge(vertexes[0], vertexes[1]))

	require.True(t, g.AddEdge(vertexes[1], vertexes[2]))
	require.True(t, g.AddEdge(vertexes[1], vertexes[3]))
	require.True(t, g.AddEdge(vertexes[1], vertexes[6]))

	require.True(t, g.AddEdge(vertexes[2], vertexes[4]))
	require.True(t, g.AddEdge(vertexes[2], vertexes[5]))
	require.True(t, g.AddEdge(vertexes[2], vertexes[6]))

	require.True(t, g.AddEdge(vertexes[3], vertexes[2]))
	require.True(t, g.AddEdge(vertexes[3], vertexes[6]))
	require.True(t, g.AddEdge(vertexes[3], vertexes[7]))

	require.True(t, g.AddEdge(vertexes[4], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[5], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[6], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[7], vertexes[8]))

	// Test Delete edge
	require.True(t, g.DelEdge(vertexes[7], vertexes[8]))
	require.False(t, g.DelEdge(vertexes[7], vertexes[8]))
}

func TestDAG_IterTopo(t *testing.T) {
	g := New()
	vertexes := []container.Int64{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// Test Add vertex
	for i := 0; i < len(vertexes); i++ {
		require.True(t, g.AddVertex(vertexes[i]))
	}

	// Test Add edge positive
	require.True(t, g.AddEdge(vertexes[0], vertexes[1]))

	require.True(t, g.AddEdge(vertexes[1], vertexes[2]))
	require.True(t, g.AddEdge(vertexes[1], vertexes[3]))
	require.True(t, g.AddEdge(vertexes[1], vertexes[6]))

	require.True(t, g.AddEdge(vertexes[2], vertexes[4]))
	require.True(t, g.AddEdge(vertexes[2], vertexes[5]))
	require.True(t, g.AddEdge(vertexes[2], vertexes[6]))

	require.True(t, g.AddEdge(vertexes[3], vertexes[2]))
	require.True(t, g.AddEdge(vertexes[3], vertexes[6]))
	require.True(t, g.AddEdge(vertexes[3], vertexes[7]))

	require.True(t, g.AddEdge(vertexes[4], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[5], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[6], vertexes[8]))

	require.True(t, g.AddEdge(vertexes[7], vertexes[8]))

	it := g.IterTopo()
	require.True(t, it.Valid())
	require.Equal(t, it.Next(), []container.Comparer{container.Int64(0)})
	require.Equal(t, it.Next(), []container.Comparer{container.Int64(1)})
	require.Equal(t, it.Next(), []container.Comparer{container.Int64(3)})
	require.Equal(t, it.Next(), []container.Comparer{container.Int64(2), container.Int64(7)})
	require.Equal(t, it.Next(), []container.Comparer{container.Int64(4), container.Int64(5), container.Int64(6)})
	require.Equal(t, it.Next(), []container.Comparer{container.Int64(8)})
	require.False(t, it.Valid())
}
