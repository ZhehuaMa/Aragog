package graphAlgorithm

import (
	"github.com/zhehuama/Aragog/graphModel"
	"testing"
)

var dumbbell = []graphModel.Edge{
	{
		U:      "0",
		V:      "1",
		Weight: 1,
	},
	{
		U:      "1",
		V:      "2",
		Weight: 1,
	},
	{
		U:      "2",
		V:      "3",
		Weight: 1,
	},
	{
		U:      "3",
		V:      "0",
		Weight: 1,
	},
	{
		U:      "4",
		V:      "5",
		Weight: 1,
	},
	{
		U:      "5",
		V:      "6",
		Weight: 1,
	},
	{
		U:      "6",
		V:      "7",
		Weight: 1,
	},
	{
		U:      "4",
		V:      "7",
		Weight: 1,
	},
	{
		U:      "0",
		V:      "4",
		Weight: 1,
	},
}

func createUndirectedGraph(edges []graphModel.Edge) graphModel.Graph {
	g := new(graphModel.UndirectedGraph)
	for _, e := range edges {
		g.AddEdge(e)
	}
	return g
}

func checkGraphSize(t *testing.T, graphs []graphModel.Graph, size int) bool {
	for i, g := range graphs {
		if len(g.GetNodes()) > size {
			t.Errorf("The size of subgraph %d should be equal to or less than %d, but %d is gotten.", i, size, len(graphs[0].GetNodes()))
			return false
		}
	}
	return true
}

func TestGettingSubGraph(t *testing.T) {
	g := createUndirectedGraph(dumbbell)

	g1, g2 := getTwoSubGraphs(g, []graphModel.Edge{dumbbell[len(dumbbell)-1]})

	if len(g1.GetNodes()) != 4 || len(g1.GetEdges()) != 4 {
		t.Errorf("Expecting a (4, 4) graph, but a (%d, %d) graph is seen.", len(g1.GetNodes()), len(g1.GetEdges()))
	}
	if len(g2.GetNodes()) != 4 || len(g2.GetEdges()) != 4 {
		t.Errorf("Expecting a (4, 4) graph, but a (%d, %d) graph is seen.", len(g2.GetNodes()), len(g2.GetEdges()))
	}
}

func TestSplitDumbbell1(t *testing.T) {
	g := createUndirectedGraph(dumbbell)

	maxSize := 4
	graphs, cut := cutToK(g, maxSize)

	if len(graphs) != 2 {
		t.Errorf("Expecting 2 subgraphs, but %d is gotten.", len(graphs))
	}
	if len(cut) != 1 {
		t.Errorf("Expecting 1 cut edge, but %d is gotten.", len(cut))
	}
	if !checkGraphSize(t, graphs, maxSize) {
		t.Errorf("Checking subgraphs' size failure.")
	}
}

func TestSplitDumbbell2(t *testing.T) {
	g := createUndirectedGraph(dumbbell)

	maxSize := 1
	graphs, cut := cutToK(g, maxSize)

	if len(graphs) != 8 {
		t.Errorf("Expecting 8 subgraphs, but %d is gotten.", len(graphs))
	}
	if len(cut) != 9 {
		t.Errorf("Expecting 9 cut edge, but %d is gotten.", len(cut))
	}
	if !checkGraphSize(t, graphs, maxSize) {
		t.Errorf("Checking subgraphs' size failure.")
	}
}

func TestSplitDumbbell3(t *testing.T) {
	g := createUndirectedGraph(dumbbell)

	maxSize := 8
	graphs, cut := cutToK(g, maxSize)

	if len(graphs) != 1 {
		t.Errorf("Expecting 1 subgraphs, but %d is gotten.", len(graphs))
	}
	if len(cut) != 0 {
		t.Errorf("Expecting 0 cut edge, but %d is gotten.", len(cut))
	}
	if !checkGraphSize(t, graphs, maxSize) {
		t.Errorf("Checking subgraphs' size failure.")
	}
}

func TestSplitDumbbell4(t *testing.T) {
	g := createUndirectedGraph(dumbbell)

	maxSize := 6
	graphs, cut := cutToK(g, maxSize)

	if len(graphs) != 2 {
		t.Errorf("Expecting 1 subgraphs, but %d is gotten.", len(graphs))
	}
	if len(cut) != 1 {
		t.Errorf("Expecting 0 cut edge, but %d is gotten.", len(cut))
	}
	if !checkGraphSize(t, graphs, maxSize) {
		t.Errorf("Checking subgraphs' size failure.")
	}
}
