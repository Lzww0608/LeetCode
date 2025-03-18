func elementInNums(nums []int, queries [][]int) []int {
    n, m := len(nums), len(queries)
    ans := make([]int, m)

    for i := 0; i < m; i++ {
        x, y := queries[i][0], queries[i][1]
        if x % (n * 2) < n {
            x %= n 
            if y >= n - x {
                ans[i] = -1
            } else {
                ans[i] = nums[x + y] 
            }
        } else {
            x %= n 
            if y >= x {
                ans[i] = -1
            } else {
                ans[i] = nums[y]
            }
        }
    }

    return ans
}