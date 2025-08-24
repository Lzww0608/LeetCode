/**
 * Definition for a rope tree node.
 * type RopeTreeNode struct {
 * 	   len   int
 * 	   val   string
 * 	   left  *RopeTreeNode
 * 	   right *RopeTreeNode
 * }
 */
func getKthCharacter(root *RopeTreeNode, k int) (ans byte) {
    var dfs func(*RopeTreeNode)
    dfs = func(root *RopeTreeNode) {
        if root == nil || k <= 0 {
            return
        }
        dfs(root.left)
        dfs(root.right)
        if root.val != "" {
            if k <= len(root.val) {
                ans = root.val[k - 1]
            }
            k -= len(root.val)
        }
    }
    dfs(root)

    return
}