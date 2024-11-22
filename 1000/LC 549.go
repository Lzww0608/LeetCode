/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) (ans int) {
    
    var dfs func(root *TreeNode) (int, int) 
    dfs = func(root *TreeNode) (asc, desc int) {
        if root == nil {
            return 0, 0
        }
        asc, desc = 1, 1
        l_asc, l_desc := dfs(root.Left)
        r_asc, r_desc := dfs(root.Right)
        if root.Left != nil {
            if root.Left.Val + 1 == root.Val {
                asc = max(asc, l_asc + 1)
                if root.Right != nil && root.Right.Val - 1 == root.Val {
                    ans = max(ans, l_asc + 1 + r_desc)
                }
            } else if root.Left.Val - 1 == root.Val {
                desc = max(desc, l_desc + 1)
                if root.Right != nil && root.Right.Val + 1 == root.Val {
                    ans = max(ans, l_desc + 1 + r_asc)
                }
            }
        } 

        if root.Right != nil {
            if root.Right.Val + 1 == root.Val {
                asc = max(asc, 1 + r_asc)
            } else if root.Right.Val - 1 == root.Val {
                desc = max(desc, 1 + r_desc)
            }
        }
        ans = max(ans, asc, desc)
        return 
    }

    dfs(root)

    return 
}