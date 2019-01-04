package data

type Stack []interface{}
type Queue []interface{}

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

type DirectedGraphNode struct {
	Val  int
	To   []*DirectedGraphNode
	From []*DirectedGraphNode
}

type Graph []*GraphNode

type DirectedGraph []*DirectedGraphNode

func (s Stack) Push(v interface{}) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, interface{}) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (q Queue) Push(v interface{}) Queue {
	return append(q, v)
}

func (q Queue) Pop() (Queue, interface{}) {
	return q[1:], q[0]
}

func visit(v *GraphNode, visited []*GraphNode) bool {
	added := false
	for i := 0; i < len(visited); i++ {
		if visited[i] == v {
			break
		}
		if visited[i] == nil {
			visited[i] = v
			added = true
			break
		}
	}
	return added
}

func (g Graph) DFS(root, dest *GraphNode) *GraphNode {
	if root == nil || root == dest {
		return root
	}
	visited := make([]*GraphNode, len(g))
	stack := Stack{}
	stack = stack.Push(root)

	for len(stack) > 0 {
		tStack, tVal := stack.Pop()
		stack = tStack
		root = tVal.(*GraphNode)
		if visit(root, visited) == true {
			if root == dest {
				return root
			}
		}
		for _, v := range root.Neighbors {
			if visit(v, visited) == true {
				stack = stack.Push(v)
			}
		}
	}

	return root
}

func (g Graph) BFS(root, dest *GraphNode) *GraphNode {
	if root == nil || root == dest {
		return root
	}
	visited := make([]*GraphNode, len(g))
	queue := Queue{}
	queue = queue.Push(root)

	for len(queue) > 0 {
		tQueue, tVal := queue.Pop()
		queue = tQueue
		root = tVal.(*GraphNode)
		if visit(root, visited) == true {
			if root == dest {
				return root
			}
		}
		for _, v := range root.Neighbors {
			if visit(v, visited) == true {
				queue = queue.Push(v)
			}
		}

	}

	return root
}
