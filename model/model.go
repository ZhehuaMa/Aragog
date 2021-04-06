package model

type Node string

type Edge struct {
	U       Node
	V       Node
	Weight  float64
	Payload interface{}
}

type Graph interface {
	GetNodes() []Node
	AddNode(u Node)
	RemoveNode(u Node)
	GetEdges() []Edge
	GetEdgesOf(u Node) ([]Edge, error)
	AddEdge(e Edge)
	RemoveEdge(u, v Node) error
	Copy() Graph
}
