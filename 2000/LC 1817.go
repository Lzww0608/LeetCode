func findingUsersActiveMinutes(logs [][]int, k int) []int {
    ans := make([]int, k)
    m := map[int]map[int]bool{}
    
    for _, log := range logs {
        a, b := log[0], log[1]
        if _, ok := m[a]; !ok {
            m[a] = make(map[int]bool)
        }
        m[a][b] = true
    }

    for _, x := range m {
        ans[len(x)-1]++
    }

    return ans
}
