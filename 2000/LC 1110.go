/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func delNodes(root *TreeNode, to_delete []int) (ans []*TreeNode) {
    m := make(map[int]bool)
	for _, x := range to_delete {
		m[x] = true
	}

	var dfs func(*TreeNode) *TreeNode 
	dfs = func(root *TreeNode) *TreeNode {
        if root == nil {
            return nil
        }

        root.Left = dfs(root.Left)
        root.Right = dfs(root.Right)
        if m[root.Val] {
            if root.Left != nil {
                ans = append(ans, root.Left)
            }
            if root.Right != nil {
                ans = append(ans, root.Right)
            }
            return nil
        }
        return root
    }
    if dfs(root) != nil {
        ans = append(ans, root)
    }

    return
 }