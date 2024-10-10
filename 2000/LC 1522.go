/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func diameter(root *Node) int {
    ans := 0

    var dfs func(*Node) int
    dfs = func(root *Node) (res int) {
        if root == nil {
            return 
        }

        a, b := 0, 0
        for _, v := range root.Children {
            cur := dfs(v)
            if cur > a {
                a, b = cur, a
            } else if cur > b {
                b = cur
            }
        }
        ans = max(ans, a + b)

        return max(a, b) + 1
    }
    dfs(root)

    return ans
}

