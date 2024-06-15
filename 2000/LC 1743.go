func restoreArray(adjacentPairs [][]int) []int {
    n := len(adjacentPairs) + 1
    edge := map[int][]int{}
    for _, v := range adjacentPairs {
        a, b := v[0], v[1]
        edge[a] = append(edge[a], b)
        edge[b] = append(edge[b], a)
    }
    ans := make([]int, n)
    for k, e := range edge {
        if len(e) == 1 {
            ans[0] = k
            break
        }
    }
    ans[1] = edge[ans[0]][0]
    for i := 2; i < n; i++ {
        v := edge[ans[i-1]]
        if ans[i-2] == v[0] {
            ans[i] = v[1]
        } else {
            ans[i] = v[0]
        }
    }
    return ans
}
