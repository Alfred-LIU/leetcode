package Tree

import (
	"strconv"
	"strings"
)

/*
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.
*/

type Codec struct {
	seperator string
	empty     string
}

func Constructor() Codec {
	return Codec{",", "#"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return this.empty
	}
	res := strconv.Itoa(root.Val)
	res += this.seperator
	res += this.serialize(root.Left)
	res += this.seperator
	res += this.serialize(root.Right)

	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	a := strings.Split(data, this.seperator)
	idx := 0
	return this.dfs297(a, &idx)
}

func (this *Codec) dfs297(in []string, idx *int) *TreeNode {
	if *idx >= len(in) {
		return nil
	}
	if in[*idx] == this.empty {
		return nil
	}

	val, _ := strconv.Atoi(in[*idx])
	node := &TreeNode{
		Val: val,
	}
	*idx = *idx + 1
	node.Left = this.dfs297(in, idx)
	*idx = *idx + 1
	node.Right = this.dfs297(in, idx)
	return node
}
