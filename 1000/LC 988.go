/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) (ans string) {
    ans = "{"
    path := []byte{}

    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        path = append(path, byte('a' + root.Val))
        if root.Left == nil && root.Right == nil {
            tmp := make([]byte, len(path))
            for i := range tmp {
                tmp[i] = path[len(path)-i-1]
            }
            ans = min(ans, string(tmp))
            path = path[:len(path)-1]
            return
        }
        if root.Left != nil {
            dfs(root.Left)
        }
        if root.Right != nil {
            dfs(root.Right)
        }
        path = path[:len(path)-1]
    }
    dfs(root)

    return
}