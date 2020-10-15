package parser

import (
	"fmt"
	"strings"
)

type Queue struct {
	Inner []*Node
}

func NewQueue() *Queue {
	return &Queue{Inner: make([]*Node, 0)}
}

func (q *Queue) Push(node *Node) {
	q.Inner = append(q.Inner, node)
}

func (q *Queue) String() string {
	nodes := make([]string, 0)
	for _, n := range q.Inner {
		nodes = append(nodes, fmt.Sprint(n.ParsedInstruction))
	}
	return strings.Join(nodes, " -> ")
}
