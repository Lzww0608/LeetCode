func findMinimumTime(strength []int, k int) int {
    n := len(strength)
    m := 1 << n 
    f := make([]int, m)
    for i := range f {
        f[i] = math.MaxInt32
    }
    f[0] = 0

    for s := 1; s < m; s++ {
        for j := 0; j < n; j++ {
            if (s >> j) & 1 == 1 {
                t := bits.OnesCount(uint(s)) - 1
                x := 1 + k * t
                f[s] = min(f[s], f[s^(1<<j)] + (strength[j] + x - 1) / x)
            }
        }
    } 

    return f[m-1]
}