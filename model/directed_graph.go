package model

import "errors"

type DirectedGraph struct {
	adjacentList map[Node]adjacentList
}

func (g *DirectedGraph) init() {
	if g.adjacentList == nil {
		g.adjacentList = make(map[Node]adjacentList)
	}
}

func (g *DirectedGraph) GetNodes() []Node {
	g.init()

	nodes := make([]Node, 0)
	for v := range g.adjacentList {
		nodes = append(nodes, v)
	}
	return nodes
}

func (g *DirectedGraph) AddNode(u Node) {
	g.init()

	if _, ok := g.adjacentList[u]; !ok {
		g.adjacentList[u] = make(adjacentList, 0)
	}
}

func (g *DirectedGraph) RemoveNode(u Node) {
	g.init()

	if _, ok := g.adjacentList[u]; !ok {
		return
	}

	for v, edges := range g.adjacentList {
		if v == u {
			continue
		}
		for i := range edges {
			if edges[i].V == u {
				edges[len(edges)-1], edges[i] = edges[i], edges[len(edges)-1]
				edges = edges[:len(edges)-1]
				g.adjacentList[v] = edges
				break
			}
		}
	}

	delete(g.adjacentList, u)
}

func (g *DirectedGraph) GetEdges() []Edge {
	g.init()

	allEdges := make([]Edge, 0)
	for _, edges := range g.adjacentList {
		for _, e := range edges {
			allEdges = append(allEdges, e)
		}
	}
	return allEdges
}

func (g *DirectedGraph) GetEdgesOf(u Node) ([]Edge, error) {
	g.init()

	if _, ok := g.adjacentList[u]; !ok {
		return nil, errors.New("Node " + string(u) + " doesn't exist.")
	}

	uEdges := g.adjacentList[u]
	edges := make([]Edge, 0)
	edges = append(edges, uEdges...)
	return edges, nil
}

func (g *DirectedGraph) GetEdgeOf(u, v Node) (Edge, error) {
	if _, ok := g.adjacentList[u]; !ok {
		return Edge{}, errors.New("Vertex " + string(u) + " doesn't exist.")
	}

	for _, e := range g.adjacentList[u] {
		if e.V == v {
			return e, nil
		}
	}

	return Edge{}, errors.New("Edge (" + string(u) + ", " + string(v) + ") doesn't exist.")
}

func (g *DirectedGraph) AddEdge(e Edge) {
	g.init()

	g.AddNode(e.U)
	g.AddNode(e.V)

	for _, edge := range g.adjacentList[e.U] {
		if edge.V == e.V {
			return
		}
	}

	g.adjacentList[e.U] = append(g.adjacentList[e.U], e)
}

func (g *DirectedGraph) RemoveEdge(u, v Node) error {
	g.init()

	var ok bool
	if _, ok = g.adjacentList[u]; !ok {
		return errors.New("Node " + string(u) + " doesn't exist.")
	}
	if _, ok = g.adjacentList[v]; !ok {
		return errors.New("Node " + string(v) + " doesn't exist.")
	}

	removeEdge(g.adjacentList, u, v)

	return nil
}

func (g *DirectedGraph) Copy() Graph {
	g.init()

	newGraph := new(DirectedGraph)
	nodes := g.GetNodes()
	for _, v := range nodes {
		newGraph.AddNode(v)
	}
	edges := g.GetEdges()
	for _, e := range edges {
		newGraph.AddEdge(e)
	}
	return newGraph
}
