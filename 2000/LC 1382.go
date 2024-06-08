/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    a := []int{}

    var dfs func(*TreeNode) 
    dfs = func(root *TreeNode)  {
        if root == nil {
            return 
        }
        dfs(root.Left)
        a = append(a, root.Val)
        dfs(root.Right)
    }
    dfs(root)

    n := len(a)
    var construct func(int, int) *TreeNode
    construct = func(l, r int) *TreeNode {
        if l > r {
            return nil
        } else if l == r {
            return &TreeNode{a[l], nil, nil}
        }
        mid := l + ((r - l) >> 1)
        root := &TreeNode{a[mid], nil, nil}
        root.Left = construct(l, mid - 1)
        root.Right = construct(mid + 1, r)
        return root
    }

    return construct(0, n - 1)
}

