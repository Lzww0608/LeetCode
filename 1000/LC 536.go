/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func str2tree(s string) *TreeNode {
    st := []*TreeNode{}
    n := len(s)
    
    for i := 0; i < n; i++ {
        if s[i] == ')' {
            st = st[:len(st)-1]
        } else if s[i] == '(' {
            continue
        } else {
            j := i 
            for j < n && s[j] != '(' && s[j] != ')' {
                j++
            }
            x, _ := strconv.Atoi(s[i:j])
            node := &TreeNode{x, nil, nil}
            if len(st) > 0 {
                if st[len(st)-1].Left == nil {
                    st[len(st)-1].Left = node
                } else {
                    st[len(st)-1].Right = node
                }
            } 
            st = append(st, node)
            i = j - 1
        }
    }

    if len(st) == 0 {
        return nil
    }
    return st[0]
}