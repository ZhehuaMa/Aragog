package algorithm

import (
	"github.com/zhehuama/Aragog/model"
	"github.com/zhehuama/Aragog/utils"
)

func bfs(g model.Graph, s, t model.Node) ([]model.Node, map[model.Node]model.Node) {
	allNodes := g.GetNodes()
	if len(allNodes) == 0 {
		return nil, nil
	}

	queue := utils.NewListQueue()
	traverseOrder := make([]model.Node, 0)
	parents := make(map[model.Node]model.Node)
	visited := make(map[model.Node]bool)

	queue.Push(s)
	parents[s] = ""
	visited[s] = true
	traverseOrder = append(traverseOrder, s)

	for !queue.IsEmpty() {
		s = queue.Pop().(model.Node)
		if t != "" && s == t {
			return traverseOrder, parents
		}
		edges, _ := g.GetEdgesOf(s)
		for _, e := range edges {
			if !visited[e.V] {
				queue.Push(e.V)
				parents[e.V] = s
				visited[e.V] = true
				traverseOrder = append(traverseOrder, e.V)
			}
		}
	}

	if t != "" {
		parents = nil
	}

	return traverseOrder, parents
}
