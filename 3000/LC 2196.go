/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    m := map[int]*TreeNode{}
    in := map[int]int{}

    for _, v := range descriptions {
        fa, ch, isLeft := v[0], v[1], v[2]
        in[ch]++
        if _, ok := m[fa]; !ok {
            m[fa] = &TreeNode{fa, nil, nil}
        }
        if _, ok := m[ch]; !ok {
            m[ch] = &TreeNode{ch, nil, nil}
        }

        if isLeft == 1 {
            m[fa].Left = m[ch]
        } else {
            m[fa].Right = m[ch]
        }
    }

    for i := range m {
        if in[i] == 0 {
            return m[i]
        }
    }

    return nil
}