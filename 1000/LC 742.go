/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findClosestLeaf(root *TreeNode, k int) int {
    g := make(map[int][]int)
    val2Node := make(map[int]*TreeNode)
    
    var dfs func(*TreeNode) 
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        val2Node[root.Val] = root
        if root.Left != nil {
            g[root.Left.Val] = append(g[root.Left.Val], root.Val)
            g[root.Val] = append(g[root.Val], root.Left.Val)
            dfs(root.Left)
        }
        if root.Right != nil {
            g[root.Right.Val] = append(g[root.Right.Val], root.Val)
            g[root.Val] = append(g[root.Val], root.Right.Val)
            dfs(root.Right)
        }
    }
    dfs(root)

    vis := make(map[int]bool)
    q := []int{k}
    vis[k] = true
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        node := val2Node[cur]
        if node.Left == nil && node.Right == nil {
            return cur
        }
        for _, x := range g[cur] {
            if !vis[x] {
                q = append(q, x)
                vis[x] = true
            }
        }
    }

    return k
}