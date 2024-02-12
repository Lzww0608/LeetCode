/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) (ans []int) {
    m := map[*TreeNode]*TreeNode{}
    var dfs func(*TreeNode, *TreeNode ,int)
    dfs = func(root *TreeNode, parent *TreeNode, depth int) {
        if root == nil || depth > k {
            return
        }
        if depth == k {
            ans = append(ans, root.Val)
            return
        }
        if root.Left != parent {
            dfs(root.Left, root, depth + 1)
        }
        if root.Right != parent {
            dfs(root.Right, root, depth + 1)
        }
        if m[root] != parent {
            dfs(m[root], root, depth + 1)
        }
    }
    
    var findParent func(*TreeNode)
    findParent = func(root *TreeNode) {
        if root == nil {
            return 
        }
        if root.Left != nil {
            m[root.Left] = root
            findParent(root.Left)
        }
        if root.Right != nil {
            m[root.Right] = root
            findParent(root.Right)
        }
    }
    findParent(root)
    dfs(target, nil, 0)

    return ans
}
