/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func inorderSuccessor(node *Node) *Node {
    if node.Right != nil {
        return leftMost(node.Right)
    }

    return findParentSuccessor(node)
}

func leftMost(node *Node) *Node {
    for node.Left != nil {
        node = node.Left
    }
    return node
}

func findParentSuccessor(node *Node) *Node {
    for node.Parent != nil && node.Parent.Right == node {
        node = node.Parent
    }
    return node.Parent
}