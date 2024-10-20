/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func boundaryOfBinaryTree(root *TreeNode) (ans []int) {
    ans = append(ans, root.Val)

    for p := root.Left; p != nil; {
        if p.Left != nil {
            ans = append(ans, p.Val)
            p = p.Left
        } else if p.Right != nil {
            ans = append(ans, p.Val)
            p = p.Right
        } else {
            break
        }
    } 
    tmp := []int{}
    for p := root.Right; p != nil; {
        if p.Right != nil {
            tmp = append(tmp, p.Val)
            p = p.Right
        } else if p.Left != nil {
            tmp = append(tmp, p.Val)
            p = p.Left
        } else {
            break
        }
    } 

    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return 
        } else if root.Left == nil && root.Right == nil {
            ans = append(ans, root.Val)
            return 
        } 
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root.Left)
    dfs(root.Right)

    if len(tmp) > 0 {
        for i := len(tmp) - 1; i >= 0; i-- {
            ans = append(ans, tmp[i])
        }
    }

    return
}