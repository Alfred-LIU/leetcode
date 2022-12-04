package Tree

/*
Given the root of a binary tree, return the length of the longest consecutive sequence path.

A consecutive sequence path is a path where the values increase by one along the path.

Note that the path can start at any node in the tree, and you cannot go from a node to its parent in the path.
*/

func longestConsecutive(root *TreeNode) int {
	res := 0

	m := make(map[*TreeNode]int)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		nodeRes := dfs(current, m)
		if nodeRes > res {
			res = nodeRes
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return res
}

func dfs(node *TreeNode, m map[*TreeNode]int) int {
	if v, ok := m[node]; ok {
		return v
	}

	res := 1
	leftRes := 0
	rightRes := 0
	if node.Left != nil && node.Val+1 == node.Left.Val {
		leftRes = dfs(node.Left, m)
	}

	if node.Right != nil && node.Val+1 == node.Right.Val {
		rightRes = dfs(node.Right, m)
	}

	if leftRes > rightRes {
		res += leftRes
	} else {
		res += rightRes
	}

	m[node] = res
	return res
}
