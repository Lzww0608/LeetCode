func matrixSum(nums [][]int) (ans int) {
    n := len(nums[0])

    for i := range nums {
        sort.Ints(nums[i])
    }

    for j := 0; j < n; j++ {
        mx := 0
        for i := range nums {
            mx = max(mx, nums[i][j])
        }
        ans += mx
    }

    return
}