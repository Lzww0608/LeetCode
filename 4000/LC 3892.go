func minOperations(nums []int, k int) int {
    n := len(nums)
    if k > n / 2 {
        return -1
    }
    cnt := 0

    for i := range nums {
        if nums[i] > nums[(i - 1 + n) % n] && nums[i] > nums[(i + 1) % n] {
            cnt++
        }
    }

    if cnt >= k {
        return 0
    }

    a := append([]int{nums[n - 1]}, nums...)
    b := append(nums, nums[0])

    return min(solve(a, k), solve(b, k))
}

func solve(a []int, k int) int {
    n := len(a)
    f := make([][]int, k + 1)
    for i := range f {
        f[i] = make([]int, n)
    }

    for i := 1; i <= k; i++ {
        f[i][i * 2 - 1] = math.MaxInt32
        for j := i * 2 - 1; j < n - 1 - (k - i) * 2; j++ {
            x := f[i][j]
            y := f[i - 1][j - 1] + max(0, max(a[j + 1], a[j - 1]) - a[j] + 1)
            f[i][j + 1] = min(x, y)
        }
    }

    return f[k][n - 1]
}