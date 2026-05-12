package trees

import (
	"reflect"
	"testing"
)

// buildTree builds a binary tree from a LeetCode-style level-order slice.
// Use -1 (or any sentinel via the nullVal arg in buildTreeWith) for nil.
// The simple form below assumes no nil placeholders in the input.
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		nodes[i] = &TreeNode{Val: v}
	}
	for i := range vals {
		l, r := 2*i+1, 2*i+2
		if l < len(vals) {
			nodes[i].Left = nodes[l]
		}
		if r < len(vals) {
			nodes[i].Right = nodes[r]
		}
	}
	return nodes[0]
}

// levelOrder returns a level-order traversal of the tree as a slice.
// Trailing nils are trimmed so complete trees round-trip cleanly.
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	out := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			continue
		}
		out = append(out, n.Val)
		queue = append(queue, n.Left, n.Right)
	}
	return out
}

func TestInvertTree(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"example1", []int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{"example2", []int{2, 1, 3}, []int{2, 3, 1}},
		{"empty", []int{}, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := levelOrder(invertTree(buildTree(tc.in)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("invertTree(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
