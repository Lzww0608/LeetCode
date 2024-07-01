/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
    builder := strings.Builder{}
    st := []*TreeNode{root}
    vis := map[*TreeNode]bool{} 
    for len(st) > 0 {
        node := st[len(st)-1]
        if vis[node] {
            if node != root {
                builder.WriteByte(')');
            }
            st = st[:len(st)-1]
        } else {
            vis[node] = true
            if node != root {
                builder.WriteByte('(')
            }
            builder.WriteString(strconv.Itoa(node.Val))
            if node.Left == nil && node.Right != nil {
                builder.WriteString("()")
            }
            if node.Right != nil {
                st = append(st, node.Right)
            }
            if node.Left != nil {
                st = append(st, node.Left)
            }
        }
    }

    return builder.String()
}
