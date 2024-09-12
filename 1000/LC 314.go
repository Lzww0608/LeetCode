import "math"
import "sort"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalOrder(root *TreeNode) (ans [][]int) {
    mn := math.MaxInt32
    type pair struct {
        x, d int
    }
    m := map[int][]pair{}

    var dfs func(*TreeNode, int, int)
    dfs = func(root *TreeNode, l, d int) {
        if root == nil {
            return 
        }

        mn = min(mn, l)
        dfs(root.Left, l - 1, d + 1)
        m[l] = append(m[l], pair{root.Val, d})
        dfs(root.Right, l + 1, d + 1)
    }

    dfs(root, 0, 0)
    for i := mn; len(m[i]) > 0; i++ {
        sort.SliceStable(m[i], func(k, j int) bool {
            return m[i][k].d < m[i][j].d 
        })
        a := []int{}
        for _, v := range m[i] {
            a = append(a, v.x)
        }
        ans = append(ans, a)
    }

    return 
}