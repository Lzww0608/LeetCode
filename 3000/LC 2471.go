/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumOperations(root *TreeNode) (ans int) {
    q := []*TreeNode{root}
    for len(q) > 0 {
        n := len(q)
        tmp := make([]int, 0, n)
        id := make([]int, n)
        for i := range id {
            id[i] = i
        }
        for t := len(q); t > 0; t-- {
            cur := q[0]
            q = q[1:]
            tmp = append(tmp, cur.Val)
            if cur.Left != nil {
                q = append(q, cur.Left)
            }
            if cur.Right != nil {
                q = append(q, cur.Right)
            }
        }
        sort.Slice(id, func(i, j int) bool {
            return tmp[id[i]] < tmp[id[j]]
        })
        ans += n
        vis := make([]bool, n)
        for _, v := range id {
            if !vis[v] {
                for !vis[v] {
                    vis[v] = true
                    v = id[v]
                }
                ans--
            }
        }
        
    }

    return 
}