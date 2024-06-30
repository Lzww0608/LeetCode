func numSquarefulPerms(nums []int) (ans int) {
    sort.Ints(nums)
    n := len(nums)
    m := map[int]bool{}

    var dfs func(int, int)
    dfs = func(idx, pre int) {
        if idx == n {
            ans++
            return 
        }
        for i := 0; i < n; i++ {
            if m[i] {
                continue
            }
            if i > 0 && nums[i] == nums[i-1] && !m[i-1] {
                continue
            }
            m[i] = true
            if idx == 0 {
                dfs(idx + 1, nums[i])
            } else if check(pre + nums[i]) {
                dfs(idx + 1, nums[i])
            }
            m[i] = false
        }
    }
    dfs(0, -1)

    return 
}

func check(x int) bool {
    sqrt := int(math.Sqrt(float64(x)))
    return sqrt * sqrt == x
}
