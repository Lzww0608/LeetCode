func minOperations(nums []int, sum int) int {
    f := make([]int, sum + 1)
    for i := range f {
        f[i] = math.MaxInt32
    }

    f[0] = 0
    for _, x := range nums {
        cur := make([]int, sum + 1)
        copy(cur, f)

        vis := make(map[int]bool)
        for l, c := x, 0; l > 0; l, c = l / 2, c + 1 {
            for r, d := l, c; r <= sum; r, d = r * 2, d + 1 {
                if vis[r] {
                    continue
                }
                vis[r] = true
                for s := 0; s + r <= sum; s++ {
                    cur[s + r] = min(cur[s + r], f[s] + d)
                }
            }
        }

        f = cur
    }

    if f[sum] >= math.MaxInt32 {
        return -1
    }
    return f[sum]
}