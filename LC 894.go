/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
    if (n & 1) == 0 {
        return nil
    }
    if n == 1 { // 基本情况
        return []*TreeNode{{0, nil, nil}}
    }
    var ans []*TreeNode
    for leftNodes := 1; leftNodes < n; leftNodes += 2 {
        leftSubtrees := allPossibleFBT(leftNodes)
        rightSubtrees := allPossibleFBT(n - 1 - leftNodes)
        for _, left := range leftSubtrees {
            for _, right := range rightSubtrees {
                root := &TreeNode{0, cloneTree(left), cloneTree(right)}
                ans = append(ans, root)
            }
        }
    }
    return ans
}

func cloneTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    newNode := &TreeNode{Val: root.Val}
    newNode.Left = cloneTree(root.Left)
    newNode.Right = cloneTree(root.Right)
    return newNode
}