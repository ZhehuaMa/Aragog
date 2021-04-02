package graphAlgorithm

import (
	"graph_cut/graphModel"
	"graph_cut/utils"
	"math"
)

/*
A simple minimum cut of undirected graph algorithm from:
STOER M, WAGNER F. A Simple Min-Cut Algorithm[J]. Journal of the ACM, 1997, 44(4): 585-591.
*/

func MinimumCut(originGraph graphModel.Graph) (float64, []graphModel.Edge) {
	g := originGraph.Copy()

	cutWeight := math.MaxFloat64
	cutEdges := make([]graphModel.Edge, 0)
	node := g.GetNodes()[0]

	mergeInfo := utils.UnionFindSet{}
	initMergeInfo(g, &mergeInfo)

	for len(g.GetNodes()) > 1 {
		weight, edges := minCutPhase(g, originGraph, node, &mergeInfo)
		if cutWeight > weight {
			cutWeight = weight
			cutEdges = edges
		}
	}

	return cutWeight, cutEdges
}

func initMergeInfo(g graphModel.Graph, mergeInfo *utils.UnionFindSet) {
	nodes := g.GetNodes()
	for _, v := range nodes {
		_ = mergeInfo.Add(string(v))
	}
}

func minCutPhase(g, originGraph graphModel.Graph, start graphModel.Node, mergeInfo *utils.UnionFindSet) (
	float64,
	[]graphModel.Edge,
) {
	s, _ := mergeInfo.Find(string(start))
	start = graphModel.Node(s)

	nodes := make([]graphModel.Node, 0)
	nodes = append(nodes, start)

	next := start
	prev := start

	for len(nodes) < len(g.GetNodes()) {
		v := getMostTightlyConnectedVertex(nodes, g)
		prev = next
		next = v
		nodes = append(nodes, v)
	}

	allEdges := originGraph.GetEdges()
	edges := make([]graphModel.Edge, 0)
	for _, e := range allEdges {
		u, _ := mergeInfo.Find(string(e.U))
		v, _ := mergeInfo.Find(string(e.V))
		if (string(next) == u && string(next) != v) || (string(next) == v && string(next) != u) {
			edges = append(edges, e)
		}
	}
	weight := float64(0)
	for _, e := range edges {
		weight += e.Weight
	}

	mergeTwoVertices(g, prev, next, mergeInfo)

	return weight, edges
}

func getMostTightlyConnectedVertex(nodes []graphModel.Node, g graphModel.Graph) graphModel.Node {
	maxWeight := float64(0)
	vertex := graphModel.Node("")
	weightOfNodes := make(map[graphModel.Node]float64)
	for _, v := range nodes {
		edges, _ := g.GetEdgesOf(v)
	search:
		for _, e := range edges {
			for _, u := range nodes {
				if e.V == u {
					continue search
				}
			}
			weightOfNodes[e.V] += e.Weight
			if weightOfNodes[e.V] > maxWeight {
				maxWeight = weightOfNodes[e.V]
				vertex = e.V
			}
		}
	}
	return vertex
}

func mergeTwoVertices(g graphModel.Graph, u, v graphModel.Node, mergeInfo *utils.UnionFindSet) {
	root, _ := mergeInfo.Unite(string(u), string(v))

	mergedEdges := make(map[graphModel.Node]float64)
	gatherEdges := func(edges []graphModel.Edge) {
		for _, e := range edges {
			mergedEdges[e.V] += e.Weight
		}
	}
	edges, _ := g.GetEdgesOf(u)
	gatherEdges(edges)
	edges, _ = g.GetEdgesOf(v)
	gatherEdges(edges)

	g.RemoveNode(u)
	g.RemoveNode(v)

	for k, weight := range mergedEdges {
		if k == u || k == v {
			continue
		}
		edge := graphModel.Edge{
			U:      graphModel.Node(root),
			V:      k,
			Weight: weight,
		}
		g.AddEdge(edge)
	}
}
