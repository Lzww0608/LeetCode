func maxIntersectionCount(y []int) (ans int) {
    m := make(map[int]int)
    n := len(y)
    for i := range y {
        y[i] <<= 1
    }

    for i := 1; i < n; i++ {
        l, r := min(y[i], y[i-1]), max(y[i], y[i-1])
        m[l]++
        m[r + 1]--
        m[y[i]]--
        m[y[i] + 1]++
    }
    m[y[n - 1]]++
    m[y[n - 1] + 1]--
    
    a := make([]int, 0, len(m))
    for k := range m {
        a = append(a, k)
    }
    sort.Ints(a)
    cur := 0
    for _, x := range a {
        cur += m[x]
        ans = max(ans, cur)
    }

    return 
}