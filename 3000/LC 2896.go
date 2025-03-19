func minOperations(s1 string, s2 string, k int) int {
    pos := []int{}
    n := len(s1)
    for i := 0; i < n; i++ {
        if s1[i] != s2[i] {
            pos = append(pos, i)
        }
    }

    if len(pos) & 1 == 1 {
        return -1
    }

    m := len(pos)
    f := make([]int, m + 1)
    for i, x := range pos {
        f[i+1] = f[i] + k
        if i - 1 >= 0 {
            f[i+1] = min(f[i+1], f[i-1] + (x - pos[i-1]) * 2)
        }
    }

    return f[m] / 2
}