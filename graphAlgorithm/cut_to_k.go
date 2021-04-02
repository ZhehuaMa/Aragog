package graphAlgorithm

import (
	"github.com/zhehuama/Aragog/graphModel"
	"github.com/zhehuama/Aragog/utils"
)

func cutToK(originGraph graphModel.Graph, maxSize int) ([]graphModel.Graph, []graphModel.Edge) {
	graphs := make([]graphModel.Graph, 0)
	edges := make([]graphModel.Edge, 0)

	queue := utils.NewListQueue()
	queue.Push(originGraph.Copy())

	for !queue.IsEmpty() {
		g := queue.Pop().(graphModel.Graph)
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

func getTwoSubGraphs(originGraph graphModel.Graph, cut []graphModel.Edge) (g1, g2 graphModel.Graph) {
	graph := originGraph.Copy()
	for _, e := range cut {
		_ = graph.RemoveEdge(e.U, e.V)
	}
	nodes1 := BFS(graph, cut[0].U)
	nodes2 := BFS(graph, cut[0].V)

	createSubGraph := func(nodes []graphModel.Node) graphModel.Graph {
		g := new(graphModel.UndirectedGraph)
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
