package algorithm

import (
	"github.com/zhehuama/Aragog/model"
	"github.com/zhehuama/Aragog/utils"
)

func cutToKUndirected(originGraph model.Graph, maxSize int) ([]model.Graph, []model.Edge) {
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
	nodes1, _ := bfs(graph, cut[0].U, "")
	nodes2, _ := bfs(graph, cut[0].V, "")

	createSubGraph := func(nodes []model.Node) model.Graph {
		g := new(model.UndirectedGraph)
		for _, v := range nodes {
			g.AddNode(v)
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

func CutToKDirected(originGraph model.Graph, maxSize int) ([]model.Graph, []model.Edge) {
	graphs := make([]model.Graph, 0)
	edges := make([]model.Edge, 0)

	graph := originGraph.Copy()
	for len(graph.GetNodes()) > maxSize {
		_, source := GetLargestComponent(graph)
		subgraph := cutDirectedGraph(graph, source, maxSize)

		var cut []model.Edge
		graph, cut = RemoveSubgraph(graph, subgraph)
		graphs = append(graphs, subgraph)
		edges = append(edges, cut...)
	}

	graphs = append(graphs, graph)

	return graphs, edges
}

func GetLargestComponent(graph model.Graph) (model.Graph, model.Node) {
	allNodes := graph.GetNodes()

	type componentType struct {
		source   model.Node
		vertices []model.Node
	}
	bfsChan := make(chan componentType)
	for s := range allNodes {
		go func(source model.Node) {
			vs, _ := bfs(graph, source, "")
			c := componentType{
				source:   source,
				vertices: vs,
			}
			bfsChan <- c
		}(s)
	}

	source := model.Node("")
	maxSize := 0
	var vertexSet []model.Node
	i := 0
	for c := range bfsChan {
		if maxSize < len(c.vertices) {
			source = c.source
			maxSize = len(c.vertices)
			vertexSet = c.vertices
		} else if maxSize == len(c.vertices) && source > c.source {
			source = c.source
			vertexSet = c.vertices
		}
		i++
		if i == len(allNodes) {
			break
		}
	}

	component := new(model.DirectedGraph)
	for _, v := range vertexSet {
		component.AddNode(v)
		edges, _ := graph.GetEdgesOf(v)
		for _, e := range edges {
			component.AddEdge(e)
		}
	}

	return component, source
}

func cutDirectedGraph(graph model.Graph, source model.Node, maxSize int) model.Graph {
	traverseOrder, _ := bfs(graph, source, "")
	if maxSize > len(traverseOrder) {
		maxSize = len(traverseOrder)
	}
	traverseOrder = traverseOrder[:maxSize]
	subVertices := make(map[model.Node]struct{})
	for _, v := range traverseOrder {
		subVertices[v] = struct{}{}
	}

	subgraph := new(model.DirectedGraph)
	for v := range subVertices {
		subgraph.AddNode(v)
		edges, _ := graph.GetEdgesOf(v)
		for _, e := range edges {
			if _, ok := subVertices[e.V]; ok {
				subgraph.AddEdge(e)
			}
		}
	}

	return subgraph
}

func RemoveSubgraph(graph, subgraph model.Graph) (model.Graph, []model.Edge) {
	vertices, subVertices := graph.GetNodes(), subgraph.GetNodes()

	cut := make([]model.Edge, 0)
	allEdges := graph.GetEdges()
	for _, e := range allEdges {
		_, ok1 := subVertices[e.U]
		_, ok2 := subVertices[e.V]
		if !ok1 && ok2 {
			cut = append(cut, e)
			continue
		}
		_, ok1 = subVertices[e.V]
		_, ok2 = subVertices[e.U]
		if !ok1 && ok2 {
			cut = append(cut, e)
			continue
		}
	}

	g := graph.Copy()
	for v := range vertices {
		if _, ok := subVertices[v]; ok {
			g.RemoveNode(v)
		}
	}

	return g, cut
}
