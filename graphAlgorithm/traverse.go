package graphAlgorithm

import (
	"github.com/zhehuama/Aragog/graphModel"
	"github.com/zhehuama/Aragog/utils"
)

func BFS(g graphModel.Graph, v graphModel.Node) []graphModel.Node {
	allNodes := g.GetNodes()
	if len(allNodes) == 0 {
		return nil
	}

	queue := utils.NewListQueue()
	traverseOrder := make([]graphModel.Node, 0)
	visited := make(map[graphModel.Node]bool)

	for _, u := range allNodes {
		visited[u] = false
	}

	queue.Push(v)

	for !queue.IsEmpty() {
		v = queue.Pop().(graphModel.Node)
		traverseOrder = append(traverseOrder, v)
		visited[v] = true
		edges, _ := g.GetEdgesOf(v)
		for _, e := range edges {
			if !visited[e.V] {
				queue.Push(e.V)
			}
		}
	}

	return traverseOrder
}
