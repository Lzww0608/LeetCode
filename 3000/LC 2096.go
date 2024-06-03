/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDirections(root *TreeNode, startValue int, destValue int) (ans string) {
    
    path := []byte{}

    var dfs func(*TreeNode, int) bool
    dfs = func(root *TreeNode, target int) bool {
        if root == nil {
            return false
        } else if root.Val == target {
            return true
        }
        path = append(path, 'L')
        if dfs(root.Left, target) {
            return true
        }
        path[len(path)-1] = 'R'
        if dfs(root.Right, target) {
            return true
        }
        path = path[:len(path)-1]
        return false
    }
    dfs(root, startValue)
    start_path := append([]byte(nil), path...)
    path = []byte{}
    dfs(root, destValue)

    i := 0
    for i < len(path) && i < len(start_path) && path[i] == start_path[i] {
        i++
    }

    return strings.Repeat("U", len(start_path[i:])) + string(path[i:])
}