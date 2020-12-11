package dag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/structs/container"
)

// InDegreeToString return strings by in-degree.
func inDegreeToString(g *DAG) string {
	s := "{ "
	it1 := g.vertexes.Iter(nil, nil)
	for it1.Valid() {
		kv1 := it1.Next()
		v1 := kv1.Value().(*Vertex)
		s += fmt.Sprintf("%v: [ ", kv1.Key())

		it2 := v1.in.Iter(nil, nil)
		for it2.Valid() {
			kv2 := it2.Next()
			v2 := kv2.Value().(*Vertex)
			_ = v2
			s += fmt.Sprintf("%v ", kv2.Key())
		}
		s += "] "
	}
	s += "}"
	return s
}

// OutDegreeToString return strings by out-degree.
func outDegreeToString(g *DAG) string {
	s := "{ "
	it1 := g.vertexes.Iter(nil, nil)
	for it1.Valid() {
		kv1 := it1.Next()
		v1 := kv1.Value().(*Vertex)
		s += fmt.Sprintf("%v: [ ", kv1.Key())

		it2 := v1.out.Iter(nil, nil)
		for it2.Valid() {
			kv2 := it2.Next()
			v2 := kv2.Value().(*Vertex)
			_ = v2
			s += fmt.Sprintf("%v ", kv2.Key())
		}
		s += "] "
	}
	s += "}"
	return s
}

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
		n, ok := kv.Value().(*Vertex)
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

	fmt.Println("outDegreeToString:", outDegreeToString(g))
	fmt.Println("inDegreeToString:", inDegreeToString(g))

	// Test Add edge negative
	require.Panics(t, func() {
		g.AddEdge(vertexes[1], vertexes[1])
	})
	//require.False(t, g.AddEdge(vertexes[1], vertexes[1]))
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
	require.NotNil(t, g.DelVertex(vertexes[3]))
	require.Nil(t, g.DelVertex(vertexes[3]))

	it1 := g.vertexes.Iter(nil, nil)
	for it1.Valid() {
		kv1 := it1.Next()
		n1 := kv1.Value().(*Vertex)
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

	fmt.Println("outDegreeToString:", outDegreeToString(g))
	fmt.Println("inDegreeToString:", inDegreeToString(g))

	var vexes []*Vertex
	it := g.IterTopo()
	require.True(t, it.Valid())

	vexes = it.Batch()
	require.Equal(t, len(vexes), 1)
	require.Equal(t, vexes[0].Key(), vertexes[0])
	require.Equal(t, vexes[0].Value(), int64(vertexes[0]*2+1))

	vexes = it.Batch()
	require.Equal(t, len(vexes), 1)
	require.Equal(t, vexes[0].Key(), vertexes[1])
	require.Equal(t, vexes[0].Value(), int64(vertexes[1]*2+1))

	vexes = it.Batch()
	require.Equal(t, len(vexes), 1)
	require.Equal(t, vexes[0].Key(), vertexes[3])
	require.Equal(t, vexes[0].Value(), int64(vertexes[3]*2+1))

	vexes = it.Batch()
	require.Equal(t, len(vexes), 2)
	require.Equal(t, vexes[0].Key(), vertexes[2])
	require.Equal(t, vexes[0].Value(), int64(vertexes[2]*2+1))
	require.Equal(t, vexes[1].Key(), vertexes[7])
	require.Equal(t, vexes[1].Value(), int64(vertexes[7]*2+1))

	vexes = it.Batch()
	require.Equal(t, len(vexes), 3)
	require.Equal(t, vexes[0].Key(), vertexes[4])
	require.Equal(t, vexes[0].Value(), int64(vertexes[4]*2+1))
	require.Equal(t, vexes[1].Key(), vertexes[5])
	require.Equal(t, vexes[1].Value(), int64(vertexes[5]*2+1))
	require.Equal(t, vexes[2].Key(), vertexes[6])
	require.Equal(t, vexes[2].Value(), int64(vertexes[6]*2+1))

	vexes = it.Batch()
	require.Equal(t, len(vexes), 1)
	require.Equal(t, vexes[0].Key(), vertexes[8])
	require.Equal(t, vexes[0].Value(), int64(vertexes[8]*2+1))

	require.False(t, it.Valid())
}
