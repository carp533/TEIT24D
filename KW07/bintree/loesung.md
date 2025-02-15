```go
func inorderTraversal(root *Node) []rune {
	// Traversiere den linken Teilbaum Inorder
	// Besuche die Wurzel
	// Traversiere den rechtenTeilbaum Inorder
	if root == nil {
		return nil
	}
	output := make([]rune, 0)
	output = append(output, inorderTraversal(root.left)...)
	output = append(output, root.data)
	output = append(output, inorderTraversal(root.right)...)
	return output
}

func preorderTraversal(root *Node) []rune {
	// Besuche die Wurzel
	// Traversiere den linken Teilbaum Preorder
	// Traversiere den rechtenTeilbaum Preorder
	if root == nil {
		return nil
	}
	output := make([]rune, 0)
	output = append(output, root.data)
	output = append(output, preorderTraversal(root.left)...)
	output = append(output, preorderTraversal(root.right)...)
	return output
}

func postorderTraversal(root *Node) []rune {
	// Traversiere den linken Teilbaum Postorder
	// Traversiere den rechtenTeilbaum Postorder
	// Besuche die Wurzel
	if root == nil {
		return nil
	}
	output := make([]rune, 0)
	output = append(output, postorderTraversal(root.left)...)
	output = append(output, postorderTraversal(root.right)...)
	output = append(output, root.data)
	return output
}
```