type TreeAncestor [][]int


func Constructor(n int, parent []int) TreeAncestor {
    m := bits.Len(uint(n))
    fa := make([][]int, n)
    for i, p := range parent {
        fa[i] = make([]int, m)
        fa[i][0] = p
    }

    for i := 0; i < m - 1; i++ {
        for x := 0; x < n; x++ {
            if p := fa[x][i]; p != -1 {
                fa[x][i+1] = fa[p][i]
            } else {
                fa[x][i+1] = -1
            }
        }
    }

    return fa
}


func (fa TreeAncestor) GetKthAncestor(node int, k int) int {
    m := bits.Len(uint(k))
    for i := 0; i < m; i++ {
        if (k >> i) & 1 > 0 {
            node = fa[node][i]
            if node < 0 {
                return -1
            }
        }
    }

    return node
}


/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
