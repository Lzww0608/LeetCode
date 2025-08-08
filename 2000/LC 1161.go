func maxLevelSum(root *TreeNode) (ans int) {
    mx := math.MinInt32
    q := []*TreeNode{root}
    d := 1
    for len(q) > 0 {
        sum := 0
        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            q = q[1:]
            sum += cur.Val
            if cur.Left != nil {
                q = append(q, cur.Left)
            }
            if cur.Right != nil {
                q = append(q, cur.Right)
            }
        }
        if sum > mx {
            mx = sum 
            ans = d
        }
        d++
    }

    return
}