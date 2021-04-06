package algorithm

import (
	"github.com/zhehuama/Aragog/model"
	"github.com/zhehuama/Aragog/utils"
)

func BFS(g model.Graph, v model.Node) []model.Node {
	allNodes := g.GetNodes()
	if len(allNodes) == 0 {
		return nil
	}

	queue := utils.NewListQueue()
	traverseOrder := make([]model.Node, 0)
	visited := make(map[model.Node]bool)

	for _, u := range allNodes {
		visited[u] = false
	}

	queue.Push(v)

	for !queue.IsEmpty() {
		v = queue.Pop().(model.Node)
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
