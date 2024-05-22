func findWinners(matches [][]int) [][]int {
    ans := make([][]int, 2)
    m := map[int]int{}
    for _, v := range matches {
        if _, ok := m[v[0]]; !ok {
            m[v[0]] = 0
        } 
        m[v[1]]++
    }

    for k, v := range m {
        if v < 2 {
            ans[v] = append(ans[v], k)
        }
    }
    
    sort.Ints(ans[0])
    sort.Ints(ans[1])

    return ans
}