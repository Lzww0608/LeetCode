/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "strconv"

func printTree(root *TreeNode) [][]string {
    if root == nil {
        return [][]string{}
    }
    h := getHeight(root)
    width := (1 << h) - 1
    ans := make([][]string, h)
    for i := range ans {
        ans[i] = make([]string, width)
    }

    var dfs func(*TreeNode, int, int)
    dfs = func(node *TreeNode, r, c int) {
        if node == nil {
            return
        }
        ans[r][c] = strconv.Itoa(node.Val)
        if h - r - 2 < 0 {
            return
        }
        offset := 1 << (h - r - 2)
        dfs(node.Left, r+1, c-offset)
        dfs(node.Right, r+1, c+offset)
    }
    dfs(root, 0, (width-1)/2)

    return ans
}

func getHeight(node *TreeNode) int {
    if node == nil {
        return 0
    }
    return max(getHeight(node.Left), getHeight(node.Right)) + 1
}


