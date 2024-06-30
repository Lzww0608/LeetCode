/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type node struct {
    row, col, val int
}
func verticalTraversal(root *TreeNode) [][]int {
    ans := [][]int{}
    nodes := []node{}
    var dfs func(*TreeNode, int, int) 
    dfs = func(root *TreeNode, row, col int) {
        if root == nil {
            return
        }
        nodes = append(nodes, node{row, col, root.Val})
        if root.Left != nil {
            dfs(root.Left, row + 1, col - 1)
        }
        if root.Right != nil {
            dfs(root.Right, row + 1, col + 1)
        }
    }
    dfs(root, 0, 0)
    sort.Slice(nodes, func(i, j int) bool {
        if nodes[i].col != nodes[j].col {
            return nodes[i].col < nodes[j].col
        } else if nodes[i].row != nodes[j].row {
            return nodes[i].row < nodes[j].row
        }
        return nodes[i].val < nodes[j].val
    })
    lastVal := -1001
    for _, node := range nodes {
        if node.col != lastVal {
            lastVal = node.col
            ans = append(ans, []int{})
        }
        ans[len(ans)-1] = append(ans[len(ans)-1], node.val)
    }
    return ans
}
