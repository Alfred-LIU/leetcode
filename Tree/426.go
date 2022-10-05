package Tree

/*
Convert a Binary Search Tree to a sorted Circular Doubly-Linked List in place.

You can think of the left and right pointers as synonymous to the predecessor and successor pointers in a doubly-linked list. For a circular doubly linked list, the predecessor of the first element is the last element, and the successor of the last element is the first element.

We want to do the transformation in place. After the transformation, the left pointer of the tree node should point to its predecessor, and the right pointer should point to its successor. You should return the pointer to the smallest element of the linked list.
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	s := Solution{}
	return s.generated(root)
}

type Solution struct {
	prev, head, tail *Node
}

func (s *Solution) generated(root *Node) *Node {
	if root == nil {
		return root
	}
	s.helper(root)
	s.head.Left = s.tail
	s.tail.Right = s.head
	return s.head
}

func (s *Solution) helper(root *Node) {
	if root == nil {
		return
	}
	s.helper(root.Left)
	if root.Left == nil && s.head == nil {
		s.head = root
	}
	if root.Right == nil {
		s.tail = root
	}

	root.Left = s.prev
	if s.prev != nil {
		s.prev.Right = root
	}
	s.prev = root
	s.helper(root.Right)
}