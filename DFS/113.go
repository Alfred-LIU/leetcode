package DFS

/*
Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int
	dfsHelper113(root, targetSum, path, &res)
	return res
}

func dfsHelper113(cur *TreeNode, targetSum int, path []int, res *[][]int) {
	if cur == nil {
		return
	}

	newTargetSum := targetSum - cur.Val
	path = append(path, cur.Val)

	if newTargetSum == 0 && cur.Left == nil && cur.Right == nil {
		cp := make([]int, len(path))
		copy(cp, path)
		*res = append(*res, cp)
		return
	}

	dfsHelper113(cur.Left, newTargetSum, path, res)
	dfsHelper113(cur.Right, newTargetSum, path, res)
}
