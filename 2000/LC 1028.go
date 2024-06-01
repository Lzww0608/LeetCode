/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(s string) *TreeNode {
    n, i := len(s), 0
    st := []*TreeNode{}
    for i < n {
        d := 0
        for s[i] == '-' {
            d++
            i++
        }
        t := 0
        for i < n && s[i] >= '0' && s[i] <= '9' {
            t = t * 10 + int(s[i] - '0')
            i++
        } 
        node := &TreeNode{t, nil, nil}
        if len(st) == 0 {
            st = append(st, node)
            continue
        }

        for len(st) > d {
            st = st[:len(st)-1]
        }
        
        if st[len(st)-1].Left != nil {
            st[len(st)-1].Right = node
        } else {
            st[len(st)-1].Left = node
        }

        st = append(st, node)
    }
    return st[0]
}
