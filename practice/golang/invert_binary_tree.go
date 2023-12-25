/**
 * Invert a binary tree
 * Example 1:
 * 	Input: root = [4,2,7,1,3,6,9]
 * 	Output: [4,7,2,9,6,3,1]
 * Example 2:
 * 	Input: root = [2,1,3]
 * 	Output: [2,3,1]
 * Example 3:
 * 	Input: root = []
 * 	Output: []
 * Constraints:
 * 	The number of nodes in the tree is in the range [0, 100].
 * 	-100 <= Node.val <= 100
**/

package problem

import (
	"fmt"
	"strings"
)

// TreeNode - Definition for a binary tree node
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func newNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func String(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var res []string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		res = append(res, fmt.Sprint(node.Val))
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return strings.Join(res, ", ")
}
