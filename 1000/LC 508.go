/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) (ans []int) {
    cnt := make(map[int]int)
    mx := 0

    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        l := dfs(root.Left)
        r := dfs(root.Right)
        sum := l + r + root.Val
        cnt[sum]++
        if cnt[sum] > mx {
            mx = cnt[sum]
            ans = []int{sum}
        } else if cnt[sum] == mx {
            ans = append(ans, sum)
        }

        return sum
    }
    dfs(root)

    return 
}