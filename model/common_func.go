package model

func removeEdge(adjacentTable map[Node]adjacentList, u, v Node) {
	edges := adjacentTable[u]
	for i := range edges {
		if edges[i].V == v {
			edges[i], edges[len(edges)-1] = edges[len(edges)-1], edges[i]
			edges = edges[:len(edges)-1]
			adjacentTable[u] = edges
			break
		}
	}
}
