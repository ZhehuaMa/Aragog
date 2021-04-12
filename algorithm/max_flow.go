package algorithm

import (
	"github.com/zhehuama/Aragog/model"
	"math"
)

type residualMatrix map[model.Node]map[model.Node]float64

func MaxFlow(originGraph model.Graph, s, t model.Node) (float64, []model.Edge) {
	matrix := getResidualMatrix(originGraph)
	graph := originGraph.Copy()

	for _, parents := bfs(graph, s, t); parents != nil; _, parents = bfs(graph, s, t) {
		path := make([]model.Node, 0)
		vertex := t
		path = append(path, vertex)
		for parents[vertex] != "" {
			vertex = parents[vertex]
			path = append(path, vertex)
		}

		minFlow := math.MaxFloat64
		for i := len(path) - 1; i > 0; i-- {
			edge, _ := graph.GetEdgeOf(path[i], path[i-1])
			if minFlow > edge.Weight {
				minFlow = edge.Weight
			}
		}

		for i := len(path) - 1; i > 0; i-- {
			u, v := path[i], path[i-1]
			matrix[u][v] -= minFlow
			matrix[v][u] += minFlow
			if math.Abs(matrix[u][v]) < 1e-5 {
				matrix[u][v] = 0
			}
		}

		graph = getGraphFromMatrix(matrix)
	}

	sourceNodes, _ := bfs(graph, s, "")
	allNodes := graph.GetNodes()

	sinkNodes := make(map[model.Node]struct{})
	for v := range allNodes {
		sinkNodes[v] = struct{}{}
	}
	for _, v := range sourceNodes {
		delete(sinkNodes, v)
	}

	cut := make([]model.Edge, 0)
	for _, v := range sourceNodes {
		edges, _ := originGraph.GetEdgesOf(v)
		for _, e := range edges {
			if _, ok := sinkNodes[e.V]; ok {
				cut = append(cut, e)
			}
		}
	}

	maxFlow := float64(0)
	for _, e := range cut {
		maxFlow += e.Weight
	}

	return maxFlow, cut
}

func getResidualMatrix(g model.Graph) residualMatrix {
	matrix := make(residualMatrix)

	edges := g.GetEdges()
	for _, e := range edges {
		if matrix[e.U] == nil {
			matrix[e.U] = make(map[model.Node]float64)
		}
		if matrix[e.V] == nil {
			matrix[e.V] = make(map[model.Node]float64)
		}
		matrix[e.U][e.V] = e.Weight
	}
	return matrix
}

func getGraphFromMatrix(m residualMatrix) model.Graph {
	g := new(model.DirectedGraph)
	for u, vertices := range m {
		for v, weight := range vertices {
			if math.Abs(weight) < 1e-5 {
				continue
			}
			edge := model.Edge{
				U:      u,
				V:      v,
				Weight: weight,
			}
			g.AddEdge(edge)
		}
	}
	return g
}
