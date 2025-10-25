func distinctAverages(nums []int) int {
    m := make(map[int]bool)

    sort.Ints(nums)
    n := len(nums)
    for i := 0; i + i < n; i++ {
        a, b := nums[i], nums[n - i - 1]
        m[a + b] = true
    }

    return len(m)
}