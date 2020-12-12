package dag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Yu-33/gohelper/structs/container"
)

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
