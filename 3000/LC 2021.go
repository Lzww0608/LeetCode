func brightestPosition(lights [][]int) int {
    m := make(map[int]int)
    for _, v := range lights {
        m[v[0] - v[1]]++
        m[v[0] + v[1] + 1]--
    }

    a := make([]int, 0, len(m))
    for k := range m {
        a = append(a, k)
    }
    sort.Ints(a)
    ans, mx := a[0], 0
    sum := 0
    for _, x := range a[1:] {
        sum += m[x]
        if sum > mx {
            mx = sum
            ans = x
        }
    }

    return ans
}