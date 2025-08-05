func maximumSubarrayXor(nums []int, queries [][]int) []int {
    n := len(nums)
    f := make([][]int, n)
    mx := make([][]int, n)
    for i := 0; i < n; i++ {
        f[i] = make([]int, n)
        mx[i] = make([]int, n)
    }

    for i := n - 1; i >= 0; i-- {
        f[i][i] = nums[i]
        mx[i][i] = nums[i]

        for j := i + 1; j < n; j++ {
            f[i][j] = f[i + 1][j] ^ f[i][j - 1]
            mx[i][j] = max(f[i][j], mx[i + 1][j], mx[i][j - 1])
        }
    }

    ans := make([]int, len(queries))
    for i, query := range queries {
        ans[i] = mx[query[0]][query[1]]
    }

    return ans
}