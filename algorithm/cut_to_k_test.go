package algorithm

import (
	"github.com/zhehuama/Aragog/model"
	"strconv"
	"testing"
)

var dumbbell = []model.Edge{
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

var component = []model.Edge{
	{
		U: "0",
		V: "1",
	},
	{
		U: "0",
		V: "3",
	},
	{
		U: "1",
		V: "2",
	},
	{
		U: "1",
		V: "3",
	},
	{
		U: "3",
		V: "2",
	},
	{
		U: "4",
		V: "3",
	},
	{
		U: "5",
		V: "2",
	},
}

func createUndirectedGraph(edges []model.Edge) model.Graph {
	g := new(model.UndirectedGraph)
	for _, e := range edges {
		g.AddEdge(e)
	}
	return g
}

func checkGraphSize(t *testing.T, graphs []model.Graph, size int) bool {
	for i, g := range graphs {
		if len(g.GetNodes()) > size || len(g.GetNodes()) == 0 {
			t.Errorf("The size of subgraph %d should be equal to or less than %d, but %d is gotten.", i, size, len(graphs[0].GetNodes()))
			return false
		}
	}
	return true
}

func generateCompleteGraph(size int, directed bool) model.Graph {
	var g model.Graph
	if directed {
		g = new(model.DirectedGraph)
	} else {
		g = new(model.UndirectedGraph)
	}

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			u := model.Node(strconv.Itoa(i))
			v := model.Node(strconv.Itoa(j))
			e := model.Edge{
				U: u,
				V: v,
				Weight: 1,
			}
			g.AddEdge(e)
			if directed {
				e.U, e.V = e.V, e.U
				g.AddEdge(e)
			}
		}
	}
	return g
}

func TestGettingSubGraph(t *testing.T) {
	g := createUndirectedGraph(dumbbell)

	g1, g2 := getTwoSubGraphs(g, []model.Edge{dumbbell[len(dumbbell)-1]})

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
	graphs, cut := cutToKUndirected(g, maxSize)

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
	graphs, cut := cutToKUndirected(g, maxSize)

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
	graphs, cut := cutToKUndirected(g, maxSize)

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
	graphs, cut := cutToKUndirected(g, maxSize)

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

func TestCompleteUndirectedGraph(t *testing.T) {
	graph := generateCompleteGraph(20, false)

	graphs, _ := cutToKUndirected(graph, 10)

	for i, g := range graphs {
		if len(g.GetNodes()) > 10 {
			t.Errorf("The size of subgraph %d is %d, expecting equal to or less than 10.", i, len(g.GetNodes()))
		}
	}
}

func TestComponent(t *testing.T) {
	graph := createDirectedGraph(component)

	c, s := getLargestComponent(graph)

	if len(c.GetNodes()) != 4 {
		t.Errorf("Expected component's size is %d, but %d gotten.", 4, len(c.GetNodes()))
	}

	if s != "0" {
		t.Errorf("Expected source node is 0, but %s gotten.", s)
	}

	edges := c.GetEdges()
	if !checkEdges(edges, component[0:5], t) {
		t.Errorf("Check component's edges failure.")
	}
}

func TestCutSubgraph(t *testing.T) {
	graph := createDirectedGraph(component)
	c, s := getLargestComponent(graph)
	subgraph := cutDirectedGraph(c, s, 4)

	if len(subgraph.GetNodes()) != 4 {
		t.Errorf("Expected subgraph's size is %d, but %d gotten.", 4, len(c.GetNodes()))
	}

	if !checkEdges(subgraph.GetEdges(), component[0:5], t) {
		t.Errorf("Check subgraph's edges failure.")
	}
}

func TestRemoveSubgraph(t *testing.T) {
	graph := createDirectedGraph(component)
	c, s := getLargestComponent(graph)
	subgraph := cutDirectedGraph(c, s, 4)
	g, edges := removeSubgraph(graph, subgraph)

	if len(g.GetNodes()) != 2 {
		t.Errorf("Expected graph size is 2, but %d gotten.", len(g.GetNodes()))
	}

	if !checkEdges(edges, component[5:7], t) {
		t.Errorf("Check graph's edges failure")
	}
}

func TestCutDirectedGraphToK(t *testing.T) {
	graph := createDirectedGraph(component)
	graphs, cut := cutToKDirected(graph, 4)

	if len(graphs) != 2 {
		t.Errorf("Expected 2 subgraphs, %d gotten.", len(graphs))
	}

	if !checkEdges(cut, component[5:7], t) {
		t.Errorf("Cut directed graph failure")
	}
}

func TestCompleteDirectedGraph(t *testing.T) {
	graph := generateCompleteGraph(100, true)

	graphs, _ := cutToKDirected(graph, 10)

	for i, g := range graphs {
		if len(g.GetNodes()) > 10 {
			t.Errorf("Subgraph %d has %d vertices, expecting 10.", i, len(g.GetNodes()))
		}
	}
}
