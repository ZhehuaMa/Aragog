package utils

import (
	"errors"
)

type nodeInfo struct {
	name string
	root string
	rank int
}
type UnionFindSet struct {
	metaInfo map[string]*nodeInfo
}

func (uf *UnionFindSet) Add(name string) error {
	if uf.metaInfo == nil {
		uf.metaInfo = make(map[string]*nodeInfo)
	}
	if _, ok := uf.metaInfo[name]; ok {
		return errors.New(name + " already exists.")
	}

	node := &nodeInfo{
		name: name,
		root: name,
		rank: 1,
	}
	uf.metaInfo[name] = node
	return nil
}

func (uf *UnionFindSet) Find(name string) (string, error) {
	if uf.metaInfo == nil {
		uf.metaInfo = make(map[string]*nodeInfo)
	}
	node, ok := uf.metaInfo[name]
	if !ok {
		return "", errors.New(name + " doesn't exist.")
	}

	if node.root == name {
		return name, nil
	}

	if root, err := uf.Find(node.root); err == nil {
		node.root = root
		return root, nil
	} else {
		return "", err
	}
}

func (uf *UnionFindSet) Unite(x, y string) (string, error) {
	if uf.metaInfo == nil {
		uf.metaInfo = make(map[string]*nodeInfo)
	}
	xRoot, err := uf.Find(x)
	if err != nil {
		return "", err
	}

	yRoot, err := uf.Find(y)
	if err != nil {
		return "", err
	}

	if xRoot == yRoot {
		return xRoot, nil
	}

	xNode, yNode := uf.metaInfo[xRoot], uf.metaInfo[yRoot]
	if xNode.rank > yNode.rank {
		yNode.root = xRoot
		return xRoot, nil
	} else if xNode.rank < yNode.rank {
		xNode.root = yRoot
		return yRoot, nil
	} else {
		yNode.root = xRoot
		xNode.rank++
		return xRoot, nil
	}
}
