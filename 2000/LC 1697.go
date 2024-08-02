func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
    m := len(queries)
    fa := make([]int, n)
    qid := make([]int, m)
    for i := range fa {
        fa[i] = i
    }
    for i := range qid {
        qid[i] = i
    }

    var find func(int) int
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[ry] = rx
        }
    }

    isConnected := func(x, y int) bool {
        return find(x) == find(y)
    }

    sort.Slice(qid, func(i, j int) bool {
        return queries[qid[i]][2] < queries[qid[j]][2]
    })
    sort.Slice(edgeList, func(i, j int) bool {
        return edgeList[i][2] < edgeList[j][2]
    })

    ans := make([]bool, m)
    i := 0
    for _, x := range qid {
        for i < len(edgeList) && edgeList[i][2] < queries[x][2] {
            merge(edgeList[i][0], edgeList[i][1])
            i++
        } 
        ans[x] = isConnected(queries[x][0], queries[x][1])
    }

    return ans
}