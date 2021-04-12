package algorithm

import (
	"github.com/zhehuama/Aragog/model"
	"math"
	"testing"
)

var diamond = []model.Edge{
	{
		U:      "0",
		V:      "1",
		Weight: 16,
	},
	{
		U:      "0",
		V:      "2",
		Weight: 13,
	},
	{
		U:      "1",
		V:      "2",
		Weight: 10,
	},
	{
		U:      "1",
		V:      "3",
		Weight: 12,
	},
	{
		U:      "2",
		V:      "1",
		Weight: 4,
	},
	{
		U:      "2",
		V:      "4",
		Weight: 14,
	},
	{
		U:      "3",
		V:      "2",
		Weight: 9,
	},
	{
		U:      "3",
		V:      "5",
		Weight: 20,
	},
	{
		U:      "4",
		V:      "3",
		Weight: 7,
	},
	{
		U:      "4",
		V:      "5",
		Weight: 4,
	},
}

var square = []model.Edge{
	{
		U:      "0",
		V:      "1",
		Weight: 3,
	},
	{
		U:      "0",
		V:      "2",
		Weight: 2,
	},
	{
		U:      "1",
		V:      "2",
		Weight: 5,
	},
	{
		U:      "1",
		V:      "3",
		Weight: 2,
	},
	{
		U:      "2",
		V:      "3",
		Weight: 3,
	},
}

func createDirectedGraph(edges []model.Edge) model.Graph {
	g := new(model.DirectedGraph)
	for _, e := range edges {
		g.AddEdge(e)
	}
	return g
}

func TestDiamond(t *testing.T) {
	graph := createDirectedGraph(diamond)

	weight, cut := MaxFlow(graph, "0", "5")

	if math.Abs(weight-23) > 1e-5 {
		t.Errorf("Expected weight is 23, but %f gotten.", weight)
	}

	if len(cut) != 3 {
		t.Errorf("Expected size of cut is 3, but %d edges gotten.", len(cut))
	}

	expectedEdges := make([]model.Edge, 3)
	expectedEdges[0] = diamond[3]
	expectedEdges[1] = diamond[8]
	expectedEdges[2] = diamond[9]

	if !checkEdges(cut, expectedEdges, t) {
		t.Errorf("Check cut failure.")
	}
}

func TestDiamond2(t *testing.T) {
	graph := createDirectedGraph(diamond)

	weight, cut := MaxFlow(graph, "1", "5")

	if math.Abs(weight-22) > 1e-5 {
		t.Errorf("Expected weight is 22, but %f gotten.", weight)
	}

	if len(cut) != 2 {
		t.Errorf("Expected size of cut is 3, but %d edges gotten.", len(cut))
	}

	expectedEdges := make([]model.Edge, 2)
	expectedEdges[0] = diamond[2]
	expectedEdges[1] = diamond[3]

	if !checkEdges(cut, expectedEdges, t) {
		t.Errorf("Check cut failure.")
	}
}

func TestSquare(t *testing.T) {
	graph := createDirectedGraph(square)

	weight, cut := MaxFlow(graph, "0", "3")

	if math.Abs(weight-5) > 1e-5 {
		t.Errorf("Expected weight is 5, but %f gotten.", weight)
	}

	if len(cut) != 2 {
		t.Errorf("Expected size of cut is 2, but %d edges gotten.", len(cut))
	}

	expectedEdges := square[0:2]

	if !checkEdges(cut, expectedEdges, t) {
		t.Errorf("Check cut failure.")
	}
}
