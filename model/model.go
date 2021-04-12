package model

type Node string

type Edge struct {
	U       Node
	V       Node
	Weight  float64
	Payload interface{}
}

type adjacentList []Edge

type Graph interface {
	GetNodes() map[Node]struct{}
	AddNode(u Node)
	RemoveNode(u Node)
	GetEdges() []Edge
	GetEdgesOf(u Node) ([]Edge, error)
	GetEdgeOf(u, v Node) (Edge, error)
	AddEdge(e Edge)
	RemoveEdge(u, v Node) error
	Copy() Graph
}
