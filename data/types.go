package data

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

type Stack []interface{}

func (s Stack) Push(v interface{}) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, interface{}) {
	l := len(s)
	return s[:l-1], s[l-1]
}

type Graph []GraphNode

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func (g Graph) DFS(root *GraphNode) *GraphNode {
	if root == nil {
		return nil
	}
	visited := make([]*GraphNode, len(g))
	stack := Stack{}
	stack = stack.Push(root)

	visit := func(v *GraphNode, visited []*GraphNode) bool {
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

	for len(stack) > 0 {
		tStack, inter := stack.Pop()
		stack = tStack
		root = inter.(*GraphNode)
		if visit(root, visited) == true {
			// Just checking for dupes/loops
		}
		for _, v := range root.Neighbors {
			if visit(v, visited) == true {
				stack = stack.Push(v)
			}
		}
	}

	return root
}
