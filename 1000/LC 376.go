func wiggleMaxLength(nums []int) int {
    n := len(nums)
    f := make([][2]int, n)
    f[0] = [2]int{1, 1}

    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            f[i][1] = f[i-1][0] + 1
            f[i][0] = f[i-1][0]
        } else if nums[i] < nums[i-1] {
            f[i][0] = f[i-1][1] + 1
            f[i][1] = f[i-1][1]
        } else {
            f[i][0] = f[i-1][0]
            f[i][1] = f[i-1][1]
        }
    }

    return max(f[n-1][0], f[n-1][1])
}