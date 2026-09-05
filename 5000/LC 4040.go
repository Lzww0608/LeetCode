func minOperations(nums []int, sum int) int {
    f := make([]int, sum + 1)
    for i := range f {
        f[i] = math.MaxInt32
    }

    f[0] = 0
    for _, x := range nums {
        cur := make([]int, sum + 1)
        copy(cur, f)

        for l, d := x, 0; l > 0; l, d = l / 2, d + 1 {
            if l > sum {
                continue
            }
            for s := 0; s + l <= sum; s++ {
                cur[s + l] = min(cur[s + l], f[s] + d)
            }
        }

        for r, d := x * 2, 1; r <= sum; r, d = r * 2, d + 1 {
            for s := 0; s + r <= sum; s++ {
                cur[s + r] = min(cur[s + r], f[s] + d)
            }
        }

        f = cur
    }

    if f[sum] >= math.MaxInt32 {
        return -1
    }
    return f[sum]
}