package algorithm

import (
	"github.com/zhehuama/Aragog/model"
	"github.com/zhehuama/Aragog/utils"
)

func cutToK(originGraph model.Graph, maxSize int) ([]model.Graph, []model.Edge) {
	graphs := make([]model.Graph, 0)
	edges := make([]model.Edge, 0)

	queue := utils.NewListQueue()
	queue.Push(originGraph.Copy())

	for !queue.IsEmpty() {
		g := queue.Pop().(model.Graph)
		if len(g.GetNodes()) <= maxSize {
			graphs = append(graphs, g)
		} else {
			_, cut := MinimumCut(g)
			edges = append(edges, cut...)
			g1, g2 := getTwoSubGraphs(g, cut)
			queue.Push(g1)
			queue.Push(g2)
		}
	}
	return graphs, edges
}

func getTwoSubGraphs(originGraph model.Graph, cut []model.Edge) (g1, g2 model.Graph) {
	graph := originGraph.Copy()
	for _, e := range cut {
		_ = graph.RemoveEdge(e.U, e.V)
	}
	nodes1 := BFS(graph, cut[0].U)
	nodes2 := BFS(graph, cut[0].V)

	createSubGraph := func(nodes []model.Node) model.Graph {
		g := new(model.UndirectedGraph)
		for _, v := range nodes {
			edges, _ := graph.GetEdgesOf(v)
			for _, e := range edges {
				g.AddEdge(e)
			}
		}
		return g
	}

	g1 = createSubGraph(nodes1)
	g2 = createSubGraph(nodes2)
	return
}
