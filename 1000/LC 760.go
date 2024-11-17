func anagramMappings(a []int, b []int) []int {
    n := len(a)
    ans := make([]int, n)
    m := map[int][]int{}
    for i, x := range b {
        m[x] = append(m[x], i)
    }

    for i, x := range a {
        ans[i] = m[x][0]
        m[x] = m[x][1:]
    }

    return ans
}