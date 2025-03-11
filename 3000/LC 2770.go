func maximumJumps(nums []int, target int) int {
    n := len(nums)
    f := make([]int, n)
    for i := range f {
        f[i] = -1
    }

    f[0] = 0
    for i := 0; i < n; i++ {
        if f[i] == -1 {
            continue
        }
        for j := i + 1; j < n; j++ {
            if abs(nums[i] - nums[j]) <= target {
                f[j] = max(f[j], f[i] + 1)
            }
        }
    }
    return f[n-1]
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}