/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() *Codec {
    return &Codec{}
}

func (c *Codec) encode(root *Node) *TreeNode {
    if root == nil {
        return nil 
    }
    cur := &TreeNode{root.Val, nil, nil}
    if len(root.Children) == 0 {
        return cur
    }
    cur.Left = c.encode(root.Children[0])
    tmp := cur.Left
    for i := 1; i < len(root.Children); i++ {
        tmp.Right = c.encode(root.Children[i])
        tmp = tmp.Right
    }
    return cur
}

func (c *Codec) decode(root *TreeNode) *Node {
    if root == nil {
        return nil 
    }
    cur := &Node{root.Val, nil}
    if root.Left == nil {
        return cur
    }

    tmp := root.Left
    for tmp != nil {
        cur.Children = append(cur.Children, c.decode(tmp))
        tmp = tmp.Right
    }
    
    return cur
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * bst := obj.encode(root);
 * ans := obj.decode(bst);
 */