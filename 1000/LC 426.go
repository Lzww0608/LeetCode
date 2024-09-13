/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

func treeToDoublyList(root *Node) *Node {
    if root == nil {
        return nil
    }
    dummy := &Node{}
    pre := dummy
    
    var dfs func(*Node)
    dfs = func(root *Node) {
        if root == nil {
            return
        }

        dfs(root.Left)
        cur := root
        cur.Left = pre
        pre.Right = cur
        pre = cur
        dfs(root.Right)

    }

    dfs(root)
    pre.Right = dummy.Right
    dummy.Right.Left = pre

    return dummy.Right
}