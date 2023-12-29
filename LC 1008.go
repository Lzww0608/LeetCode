/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    var dfs func( int, int) *TreeNode
    dfs = func(l int, r int) *TreeNode {
        root := &TreeNode{Val: preorder[l]}
        k := l+1
        for k <= r && preorder[k] < preorder[l] {
            k++
        }
        if k <= r {
            root.Right = dfs(k,r)
        } 
        if k-1 >= l+1 {
            root.Left = dfs(l+1,k-1)
        }
        return root
    }
    return dfs(0, len(preorder)-1)
}