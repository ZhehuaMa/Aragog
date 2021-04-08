package model

import "errors"

type UndirectedGraph struct {
	adjacentList map[Node]adjacentList
}

func (g *UndirectedGraph) init() {
	if g.adjacentList == nil {
		g.adjacentList = make(map[Node]adjacentList)
	}
}

func (g *UndirectedGraph) GetNodes() []Node {
	g.init()

	nodes := make([]Node, 0)
	for v := range g.adjacentList {
		nodes = append(nodes, v)
	}
	return nodes
}

func (g *UndirectedGraph) AddNode(u Node) {
	g.init()

	if _, ok := g.adjacentList[u]; !ok {
		g.adjacentList[u] = make(adjacentList, 0)
	}
}

func (g *UndirectedGraph) RemoveNode(u Node) {
	g.init()

	if _, ok := g.adjacentList[u]; !ok {
		return
	}

	edges := g.adjacentList[u]
	for _, edge := range edges {
		vEdges := g.adjacentList[edge.V]
		for i := range vEdges {
			if vEdges[i].V == u {
				vEdges[len(vEdges)-1], vEdges[i] = vEdges[i], vEdges[len(vEdges)-1]
				vEdges = vEdges[:len(vEdges)-1]
				g.adjacentList[edge.V] = vEdges
				break
			}
		}
	}

	delete(g.adjacentList, u)
}

func (g *UndirectedGraph) GetEdges() []Edge {
	g.init()

	allEdges := make([]Edge, 0)
	processedVertices := make(map[Node]struct{})
	for v, edges := range g.adjacentList {
		for _, e := range edges {
			if _, ok := processedVertices[e.V]; ok {
				continue
			}
			allEdges = append(allEdges, e)
		}
		processedVertices[v] = struct{}{}
	}
	return allEdges
}

func (g *UndirectedGraph) GetEdgesOf(u Node) ([]Edge, error) {
	g.init()

	if _, ok := g.adjacentList[u]; !ok {
		return nil, errors.New("Node " + string(u) + " doesn't exist.")
	}

	uEdges := g.adjacentList[u]
	edges := make([]Edge, 0)
	edges = append(edges, uEdges...)
	return edges, nil
}

func (g *UndirectedGraph) GetEdgeOf(u, v Node) (Edge, error) {
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

func (g *UndirectedGraph) AddEdge(e Edge) {
	g.init()

	g.AddNode(e.U)
	g.AddNode(e.V)

	for _, edge := range g.adjacentList[e.U] {
		if edge.V == e.V {
			return
		}
	}

	g.adjacentList[e.U] = append(g.adjacentList[e.U], e)
	e.U, e.V = e.V, e.U
	g.adjacentList[e.U] = append(g.adjacentList[e.U], e)
}

func (g *UndirectedGraph) RemoveEdge(u, v Node) error {
	g.init()

	var ok bool
	if _, ok = g.adjacentList[u]; !ok {
		return errors.New("Node " + string(u) + " doesn't exist.")
	}
	if _, ok = g.adjacentList[v]; !ok {
		return errors.New("Node " + string(v) + " doesn't exist.")
	}

	removeEdge(g.adjacentList, u, v)
	removeEdge(g.adjacentList, v, u)

	return nil
}

func (g *UndirectedGraph) Copy() Graph {
	g.init()

	newGraph := new(UndirectedGraph)
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
