func widestPairOfIndices(a []int, b []int) (ans int) {
    // preA[j] - preA[i] == preB[j] - preB[i]
    // preA[j] - preB[j] == preA[i] - preB[i]
    preA, preB := 0, 0
    m := make(map[int]int)
    m[0] = -1
    for i := range a {
        preA += a[i]
        preB += b[i]
        cur := preA - preB
        if v, ok := m[cur]; ok {
            ans = max(ans, i - v)
        } else {
            m[cur] = i
        }
    } 

    return
}