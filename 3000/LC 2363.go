func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    m := make(map[int]int)
    for _, v := range items1 {
        m[v[0]] += v[1]
    }
    for _, v := range items2 {
        m[v[0]] += v[1]
    }

    ans := make([][]int, 0, len(m))
    for k, v := range m {
        ans = append(ans, []int{k, v})
    }
    sort.Slice(ans, func(i, j int) bool {
        return ans[i][0] < ans[j][0]
    })

    return ans
}