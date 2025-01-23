func substringXorQueries(s string, queries [][]int) [][]int {
    n := len(s)
    cnt := make(map[int][]int)
    for d := 1; d <= min(n, 31); d++ {
        for i := 0; i + d <= n; i++ {
            x, _ := strconv.ParseInt(s[i:i+d], 2, 64)
            if _, ok := cnt[int(x)]; !ok {
                cnt[int(x)] = []int{i, i + d - 1}
            }
        }
    }

    ans := make([][]int, len(queries))
    for i, q := range queries {
        x := q[0] ^ q[1]
        if v, ok := cnt[x]; ok {
            ans[i] = v
        } else {
            ans[i] = []int{-1, -1}
        }
        
    }

    return ans
}