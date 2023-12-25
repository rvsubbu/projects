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
	"gotest.tools/assert"
	"testing"
)

func TestNilTree(t *testing.T) {
	var root *TreeNode
	expected := ""
	testHelper1(t, root, expected)
}

func Test2Level(t *testing.T) {
	node3 := newNode(3, nil, nil)
	node1 := newNode(1, nil, nil)
	root := newNode(2, node1, node3)
	expected := "2, 3, 1"
	testHelper1(t, root, expected)
}

func Test3Level(t *testing.T) {
	node9 := newNode(9, nil, nil)
	node6 := newNode(6, nil, nil)
	node3 := newNode(3, nil, nil)
	node1 := newNode(1, nil, nil)
	node7 := newNode(7, node6, node9)
	node2 := newNode(2, node1, node3)
	root := newNode(4, node2, node7)

	expected := "4, 7, 2, 9, 6, 3, 1"
	testHelper1(t, root, expected)
}

func testHelper1(t *testing.T, root *TreeNode, expected string) {
	invertedTree := invertTree(root)
	assert.Equal(t, String(invertedTree), expected)
}
