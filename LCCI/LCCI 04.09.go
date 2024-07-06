/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func BSTSequences(root *TreeNode) (ans [][]int) {
    if root == nil {
        return [][]int{[]int{}}
    }
    q := []*TreeNode{root}
    path := []int{}

    var dfs func() 
    dfs = func() {
        if len(q) == 0 {
            ans = append(ans, append([]int(nil), path...))
            return
        }

        for i := len(q); i > 0; i-- {
            node := q[0]
            q = q[1:]

            path = append(path, node.Val)

            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }

            dfs()
            if node.Left != nil {
                q = q[:len(q)-1]
            }
            if node.Right != nil {
                q = q[:len(q)-1]
            }

            q = append(q, node)
            path = path[:len(path)-1]
        }
    }
    dfs()

    return
}