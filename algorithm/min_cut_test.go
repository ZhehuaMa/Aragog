package algorithm

import (
	"github.com/zhehuama/Aragog/model"
	"math"
	"testing"
)

var line = []model.Edge{
	{
		U:      "0",
		V:      "1",
		Weight: 1,
	},
}

var angle = []model.Edge{
	{
		U:      "0",
		V:      "1",
		Weight: 2,
	},
	{
		U:      "2",
		V:      "1",
		Weight: 4,
	},
}

var triangle = []model.Edge{
	{
		U:      "0",
		V:      "1",
		Weight: 1,
	},
	{
		U:      "1",
		V:      "2",
		Weight: 2,
	},
	{
		U:      "0",
		V:      "2",
		Weight: 4,
	},
}

var exampleInPaper = []model.Edge{
	{
		U:      "1",
		V:      "2",
		Weight: 2,
	},
	{
		U:      "1",
		V:      "5",
		Weight: 3,
	},
	{
		U:      "2",
		V:      "3",
		Weight: 3,
	},
	{
		U:      "2",
		V:      "5",
		Weight: 2,
	},
	{
		U:      "2",
		V:      "6",
		Weight: 2,
	},
	{
		U:      "3",
		V:      "4",
		Weight: 4,
	},
	{
		U:      "3",
		V:      "7",
		Weight: 2,
	},
	{
		U:      "4",
		V:      "7",
		Weight: 2,
	},
	{
		U:      "4",
		V:      "8",
		Weight: 2,
	},
	{
		U:      "5",
		V:      "6",
		Weight: 3,
	},
	{
		U:      "6",
		V:      "7",
		Weight: 1,
	},
	{
		U:      "7",
		V:      "8",
		Weight: 3,
	},
}

func checkEdge(e1, e2 model.Edge) bool {
	return ((e1.U == e2.U && e1.V == e2.V) || (e1.U == e2.V && e1.V == e2.U)) && (math.Abs(e1.Weight-e2.Weight) < 1e-5)
}

func checkEdges(edges1, edges2 []model.Edge, t *testing.T) bool {
	for _, e1 := range edges1 {
		match := false
		for _, e2 := range edges2 {
			if checkEdge(e1, e2) {
				match = true
				break
			}
		}
		if !match {
			t.Errorf("Unmatched edge %v.", e1)
			return false
		}
	}
	return true
}

func TestOneNodeGraph(t *testing.T) {
	graph := new(model.UndirectedGraph)
	graph.AddNode("node")

	weight, cut := MinimumCut(graph)
	if weight != math.MaxFloat64 {
		t.Errorf("Weight should be %.2f, but %.2f seen.", math.MaxFloat64, weight)
	}
	if len(cut) != 0 {
		t.Errorf("Length of cut should be 0, but %d gotten.", len(cut))
	}
}

func TestLineCut(t *testing.T) {
	graph := createUndirectedGraph(line)

	weight, cut := MinimumCut(graph)
	if weight != 1 {
		t.Errorf("Weight should be 1, but %.2f seen.", weight)
	}
	if len(cut) != 1 {
		t.Fatalf("Length of cut should be 1, but %d gotten.", len(cut))
	}
	if !checkEdge(cut[0], line[0]) {
		t.Errorf("Unexpected edge:\nexpect %v, got %v.", line, cut[0])
	}
}

func TestAngleCut(t *testing.T) {
	graph := createUndirectedGraph(angle)

	weight, cut := MinimumCut(graph)
	if weight != 2 {
		t.Errorf("Weight should be 2, but %.2f seen.", weight)
	}
	if len(cut) != 1 {
		t.Fatalf("Length of cut should be 1, but %d gotten.", len(cut))
	}
	if !checkEdges(cut, angle[:1], t) {
		t.Error("Wrong cut.")
	}
}

func TestTriangleCut(t *testing.T) {
	graph := createUndirectedGraph(triangle)

	weight, cut := MinimumCut(graph)
	if weight != 3 {
		t.Errorf("Weight should be 3, but %.2f seen.", weight)
	}
	if len(cut) != 2 {
		t.Fatalf("Length of cut should be 2, but %d gotten.", len(cut))
	}
	if !checkEdges(cut, triangle[:2], t) {
		t.Error("Wrong cut.")
	}
}

func TestExampleCut(t *testing.T) {
	graph := createUndirectedGraph(exampleInPaper)

	weight, cut := MinimumCut(graph)
	if weight != 4 {
		t.Errorf("Weight should be 4, but %.2f seen.", weight)
	}
	if len(cut) != 2 {
		t.Fatalf("Length of cut should be 2, but %d gotten.", len(cut))
	}
	expectedCut := make([]model.Edge, 2)
	expectedCut[0] = exampleInPaper[2]
	expectedCut[1] = exampleInPaper[10]
	if !checkEdges(cut, expectedCut, t) {
		t.Error("Wrong cut.")
	}
}
