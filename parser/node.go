package parser

import "fmt"

type Node struct {
	Parent      *Node
	Instruction []string
	ParsedInstruction Instruction
	Children    []*Node
}

func newNode(instruction []string) *Node {
	return &Node{Instruction: instruction, Children: make([]*Node, 0)}
}

func (parent *Node) AddChild(children *Node) {
	children.Parent = parent
	parent.Children = append(parent.Children, children)
}

func (n *Node) Queue() *Queue {
	q := NewQueue()
	n.populateQueue(q)
	return q
}

func (n *Node) populateQueue(queue *Queue) {
	queue.Push(n)

	for _, child := range n.Children {
		child.populateQueue(queue)
	}
}

func (node *Node) Back(i int) *Node {
	for x := 0; x < i; x++ {
		if node.Parent == nil {
			fmt.Println("BAD PARENT IS NIL", node.Instruction, i)
		}
		node = node.Parent
	}
	return node
}
