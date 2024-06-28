/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    p, q := &TreeNode{math.MinInt32, nil, nil}, &TreeNode{math.MinInt32, nil, nil}
    last := &TreeNode{math.MinInt32, nil, nil}
    
    var dfs func(*TreeNode)
    dfs = func(cur *TreeNode) {
        if cur == nil {
            return 
        }
        dfs(cur.Left)
        if cur.Val < last.Val {
            if p.Val == math.MinInt32 {
                p = last
            } 
            q = cur
        }
        last = cur
        dfs(cur.Right)
    }
    dfs(root)
    p.Val, q.Val = q.Val, p.Val

    return 
}
