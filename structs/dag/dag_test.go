package dag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/structs/container"
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
		require.True(t, g.AddVertex(vertexes[i], int64(vertexes[i]*2+1)))
		require.False(t, g.AddVertex(vertexes[i], int64(vertexes[i]*2+1)))
	}

	// test internal element
	it1 := g.vertexes.Iter(nil, nil)
	for it1.Valid() {
		kv := it1.Next()
		require.NotNil(t, kv)
		require.NotNil(t, kv.Key())
		require.NotNil(t, kv.Value())
		n, ok := kv.Value().(*Node)
		require.True(t, ok)
		require.NotNil(t, n.value)
		require.NotNil(t, n.in)
		require.NotNil(t, n.out)
	}
}

func TestDAG_AddEdge(t *testing.T) {
	g := New()
	vertexes := []container.Int64{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// Test Add vertex
	for i := 0; i < len(vertexes); i++ {
		require.True(t, g.AddVertex(vertexes[i], int64(vertexes[i]*2+1)))
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
		require.True(t, g.AddVertex(vertexes[i], int64(vertexes[i]*2+1)))
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
		kv1 := it1.Next()
		n1 := kv1.Value().(*Node)
		require.NotEqual(t, kv1.Key(), vertexes[3])

		it2 := n1.in.Iter(nil, nil)
		for it2.Valid() {
			kv2 := it2.Next()
			require.NotEqual(t, kv2.Key(), vertexes[3])
		}
		it3 := n1.out.Iter(nil, nil)
		for it3.Valid() {
			kv3 := it3.Next()
			require.NotEqual(t, kv3.Key(), vertexes[3])
		}
	}
}

func TestDAG_DelEdge(t *testing.T) {
	g := New()
	vertexes := []container.Int64{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// Test Add vertex
	for i := 0; i < len(vertexes); i++ {
		require.True(t, g.AddVertex(vertexes[i], int64(vertexes[i]*2+1)))
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
		require.True(t, g.AddVertex(vertexes[i], int64(vertexes[i]*2+1)))
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

	var kvs []KV
	it := g.IterTopo()
	require.True(t, it.Valid())

	kvs = it.Next()
	require.Equal(t, len(kvs), 1)
	require.Equal(t, kvs[0].Key(), vertexes[0])
	require.Equal(t, kvs[0].Value().(*Node).value, int64(vertexes[0]*2+1))

	kvs = it.Next()
	require.Equal(t, len(kvs), 1)
	require.Equal(t, kvs[0].Key(), vertexes[1])
	require.Equal(t, kvs[0].Value().(*Node).value, int64(vertexes[1]*2+1))

	kvs = it.Next()
	require.Equal(t, len(kvs), 1)
	require.Equal(t, kvs[0].Key(), vertexes[3])
	require.Equal(t, kvs[0].Value().(*Node).value, int64(vertexes[3]*2+1))

	kvs = it.Next()
	require.Equal(t, len(kvs), 2)
	require.Equal(t, kvs[0].Key(), vertexes[2])
	require.Equal(t, kvs[0].Value().(*Node).value, int64(vertexes[2]*2+1))
	require.Equal(t, kvs[1].Key(), vertexes[7])
	require.Equal(t, kvs[1].Value().(*Node).value, int64(vertexes[7]*2+1))

	kvs = it.Next()
	require.Equal(t, len(kvs), 3)
	require.Equal(t, kvs[0].Key(), vertexes[4])
	require.Equal(t, kvs[0].Value().(*Node).value, int64(vertexes[4]*2+1))
	require.Equal(t, kvs[1].Key(), vertexes[5])
	require.Equal(t, kvs[1].Value().(*Node).value, int64(vertexes[5]*2+1))
	require.Equal(t, kvs[2].Key(), vertexes[6])
	require.Equal(t, kvs[2].Value().(*Node).value, int64(vertexes[6]*2+1))

	kvs = it.Next()
	require.Equal(t, len(kvs), 1)
	require.Equal(t, kvs[0].Key(), vertexes[8])
	require.Equal(t, kvs[0].Value().(*Node).value, int64(vertexes[8]*2+1))

	require.False(t, it.Valid())
}
