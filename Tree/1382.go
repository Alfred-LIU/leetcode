package Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sort(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	sort(root.Left, arr)
	*arr = append(*arr, root.Val)
	sort(root.Right, arr)
}

func createBalancedBST(A []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	m := (l + r) / 2
	root := &TreeNode{Val: A[m]}

	root.Left = createBalancedBST(A, l, m-1)
	root.Right = createBalancedBST(A, m+1, r)

	return root
}

func balanceBST(root *TreeNode) *TreeNode {
	var arr []int
	sort(root, &arr)
	res := createBalancedBST(arr, 0, len(arr))
	return res
}
