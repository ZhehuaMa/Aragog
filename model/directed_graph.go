package model

import "errors"

type DirectedGraph struct {
	adjacentTable map[Node]adjacentList
}

func (g *DirectedGraph) init() {
	if g.adjacentTable == nil {
		g.adjacentTable = make(map[Node]adjacentList)
	}
}

func (g *DirectedGraph) GetNodes() map[Node]struct{} {
	g.init()

	nodes := make(map[Node]struct{})
	for v := range g.adjacentTable {
		nodes[v] = struct{}{}
	}
	return nodes
}

func (g *DirectedGraph) AddNode(u Node) {
	g.init()

	if _, ok := g.adjacentTable[u]; !ok {
		g.adjacentTable[u] = make(adjacentList, 0)
	}
}

func (g *DirectedGraph) RemoveNode(u Node) {
	g.init()

	if _, ok := g.adjacentTable[u]; !ok {
		return
	}

	for v, edges := range g.adjacentTable {
		if v == u {
			continue
		}
		for i := range edges {
			if edges[i].V == u {
				edges[len(edges)-1], edges[i] = edges[i], edges[len(edges)-1]
				edges = edges[:len(edges)-1]
				g.adjacentTable[v] = edges
				break
			}
		}
	}

	delete(g.adjacentTable, u)
}

func (g *DirectedGraph) GetEdges() []Edge {
	g.init()

	allEdges := make([]Edge, 0)
	for _, edges := range g.adjacentTable {
		for _, e := range edges {
			allEdges = append(allEdges, e)
		}
	}
	return allEdges
}

func (g *DirectedGraph) GetEdgesOf(u Node) ([]Edge, error) {
	g.init()

	if _, ok := g.adjacentTable[u]; !ok {
		return nil, errors.New("Node " + string(u) + " doesn't exist.")
	}

	uEdges := g.adjacentTable[u]
	edges := make([]Edge, 0)
	edges = append(edges, uEdges...)
	return edges, nil
}

func (g *DirectedGraph) GetEdgeOf(u, v Node) (Edge, error) {
	if _, ok := g.adjacentTable[u]; !ok {
		return Edge{}, errors.New("Vertex " + string(u) + " doesn't exist.")
	}

	for _, e := range g.adjacentTable[u] {
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

	for _, edge := range g.adjacentTable[e.U] {
		if edge.V == e.V {
			return
		}
	}

	g.adjacentTable[e.U] = append(g.adjacentTable[e.U], e)
}

func (g *DirectedGraph) RemoveEdge(u, v Node) error {
	g.init()

	var ok bool
	if _, ok = g.adjacentTable[u]; !ok {
		return errors.New("Node " + string(u) + " doesn't exist.")
	}
	if _, ok = g.adjacentTable[v]; !ok {
		return errors.New("Node " + string(v) + " doesn't exist.")
	}

	removeEdge(g.adjacentTable, u, v)

	return nil
}

func (g *DirectedGraph) Copy() Graph {
	g.init()

	newGraph := new(DirectedGraph)
	nodes := g.GetNodes()
	for v := range nodes {
		newGraph.AddNode(v)
	}
	edges := g.GetEdges()
	for _, e := range edges {
		newGraph.AddEdge(e)
	}
	return newGraph
}
