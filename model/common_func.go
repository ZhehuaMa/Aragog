package model

func removeEdge(adjacentTable map[Node]adjacentList, u, v Node) {
	edges := adjacentTable[u]
	delete(edges, v)
}
