package utils

import "container/list"

type Queue interface {
	Push(v interface{})
	Front() interface{}
	Pop() interface{}
	IsEmpty() bool
}

type ListQueue struct {
	queue *list.List
}

func NewListQueue() *ListQueue {
	q := new(ListQueue)
	q.queue = list.New()
	return q
}

func (l *ListQueue) Push(v interface{}) {
	l.queue.PushBack(v)
}

func (l *ListQueue) Front() interface{} {
	e := l.queue.Front()
	return e.Value
}

func (l *ListQueue) Pop() interface{} {
	e := l.queue.Front()
	v := l.queue.Remove(e)
	return v
}

func (l *ListQueue) IsEmpty() bool {
	return l.queue.Len() == 0
}
