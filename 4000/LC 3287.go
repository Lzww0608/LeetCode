func maxValue(nums []int, k int) (ans int) {
    a := find(nums, k)
    n := len(nums)
    tmp := make([]int, n)
    copy(tmp, nums)
    for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
        tmp[i], tmp[j] = tmp[j], tmp[i]
    }
    b := find(tmp, k)
    for i := k - 1; i < n - k; i++ {
        for x := range a[i] {
            for y := range b[n-i-2] {
                ans = max(ans, x ^ y)
            }
        }
    }

    return
}

func find(a []int, k int) (ans []map[int]bool) {
    f := make([]map[int]bool, k + 1)
    for i := range f {
        f[i] = make(map[int]bool)
    }
    f[0][0] = true
    for i := range a {
        for j := min(k - 1, i + 1); j >= 0; j-- {
            for x := range f[j] {
                f[j+1][x|a[i]] = true
            }
        }
        cur := make(map[int]bool)
        for x := range f[k] {
            cur[x] = true
        }
        ans = append(ans, cur)
    }
    return
}