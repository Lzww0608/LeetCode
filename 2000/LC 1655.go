func canDistribute(nums []int, quantity []int) bool {
    m := len(quantity)
    s := make([]int, 1 << m)
    for i := 0; i < m; i++ {
        k := 1 << i
        for j := 0; j < k; j++ {
            s[j|k] = s[j] + quantity[i]
        }
    }

    cnt := make(map[int]int)
    for _, x := range nums {
        cnt[x]++
    }
    n := len(cnt)

    f := make([][]bool, n + 1)
    for i := range f {
        f[i] = make([]bool, 1 << m)
        f[i][0] = true
    }

    i := 0
    for _, v := range cnt {
        for j := range f[i] {
            if f[i][j] {
                f[i+1][j] = true
            } else {
                for k := j; k > 0; k = (k - 1) & j{
                    if s[k] <= v && f[i][k ^ j] {
                        f[i+1][j] = true
                        break
                    }
                }
            }
        } 
        i++
    }

    return f[n][(1<<m)-1]
}