func hIndex(citations []int) int {
    m := map[int]int{}
    mx := 0
    for _, x := range citations {
        m[x]++
        mx = max(mx, x)
    }

    for h := mx; h > 0; h-- {
        m[h] += m[h+1] 
        if m[h] >= h {
            return h
        }
    }
    return 0
}
