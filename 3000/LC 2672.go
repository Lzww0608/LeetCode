func colorTheArray(n int, queries [][]int) []int {
    m := len(queries)
    ans := make([]int, m)
    cnt := 0

    f := make([]int, n + 2)
    for i, v := range queries {
        j := v[0]
        if f[j + 1] != 0 {
            if f[j + 1] == f[j + 2] {
                cnt--
            }
            if f[j + 1] == f[j] {
                cnt--
            }
        }
        
        f[j + 1] = v[1]
        if f[j + 1] == f[j + 2] {
            cnt++
        }
        if f[j + 1] == f[j] {
            cnt++
        }
        
        ans[i] = cnt
    }

    return ans
}