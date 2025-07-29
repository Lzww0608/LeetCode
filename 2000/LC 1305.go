/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    a := []int{}
    dfs(root1, &a)

    b := []int{}
    dfs(root2, &b)

    m, n := len(a), len(b)
    ans := make([]int, 0, m + n)
    for i, j := 0, 0; i < m || j < n; {
        if j >= n || i < m && a[i] <= b[j] {
            ans = append(ans, a[i])
            i++
        } else {
            ans = append(ans, b[j])
            j++
        }
    }

    return ans
}

func dfs(root *TreeNode, a *[]int) {
    if root == nil {
        return 
    }
    dfs(root.Left, a)
    
    *a = append(*a, root.Val)

    dfs(root.Right, a)
}