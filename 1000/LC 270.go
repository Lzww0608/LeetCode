/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    ans := root.Val
    for root != nil {
        if float64(root.Val) == target {
            return root.Val
        }
        
        x, y := math.Abs(float64(root.Val) - target), math.Abs(float64(ans) - target)
        if x < y || x == y && root.Val < ans {
            ans = root.Val
        }

        if target < float64(root.Val) {
            root = root.Left
        } else {
            root = root.Right
        }
    }

    return ans
}
 