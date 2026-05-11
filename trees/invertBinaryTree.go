package trees

func invertTree(root *TreeNode) *TreeNode {

	// Base case : If the current node is nil, return nil
	if root == nil {
		return root
	}

	// Recursively invert left and right subtrees and swap them
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	// Return the updated root node
	return root
}