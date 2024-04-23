package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		node := TreeNode{Data: data}
		return &node
	} else {
		// root.Parent = root
		if data < root.Data {
			root.Left = BTreeInsertData(root.Left, data)
			root.Parent = root
		} else {
			root.Right = BTreeInsertData(root.Right, data)
			root.Parent = root
		}
	}
	return root
}
